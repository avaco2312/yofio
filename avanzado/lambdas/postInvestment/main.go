package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"yofio/asigna"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/segmentio/ksuid"
)

type Request struct {
	Investment int32
}

func insert(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := Request{}
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil || req.Investment <= 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"id": "0", "error": "petición incorrecta"}`,
		}, nil
	}
	ncredito := asigna.NewCreditAssigner()
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		}, nil
	}
	svc := dynamodb.NewFromConfig(cfg)
	id := ksuid.New().String()
	x, y, z, err := ncredito.Assign(req.Investment)
	if err != nil {
		_, err1 := svc.TransactWriteItems(context.TODO(), &dynamodb.TransactWriteItemsInput{
			TransactItems: []types.TransactWriteItem{
				{
					Put: &types.Put{
						Item: map[string]types.AttributeValue{
							"Id":       &types.AttributeValueMemberS{Value: id},
							"ANoExImp": &types.AttributeValueMemberN{Value: fmt.Sprint(req.Investment)},
						},
						TableName: aws.String("yofio"),
					},
				},
				{
					Update: &types.Update{
						ExpressionAttributeNames: map[string]string{
							"#anoexcan": "ANoExCan",
							"#anoeximp": "ANoExImp",
						},
						ExpressionAttributeValues: map[string]types.AttributeValue{
							":investment": &types.AttributeValueMemberN{Value: fmt.Sprint(req.Investment)},
							":uno":        &types.AttributeValueMemberN{Value: "1"},
						},
						Key: map[string]types.AttributeValue{
							"Id": &types.AttributeValueMemberS{Value: "*"},
						},
						TableName:        aws.String("yofio"),
						UpdateExpression: aws.String("SET #anoexcan = #anoexcan + :uno, #anoeximp = #anoeximp + :investment"),
					},
				},
			},
		})
		if err1 != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       fmt.Sprintf(`{"error": "%s"}`, err1.Error()),
			}, nil
		} else {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       fmt.Sprintf(`{"id": "%s", "error": "%s"}`, id, err.Error()),
			}, nil
		}
	}
	_, err = svc.TransactWriteItems(context.TODO(), &dynamodb.TransactWriteItemsInput{
		TransactItems: []types.TransactWriteItem{
			{
				Put: &types.Put{
					Item: map[string]types.AttributeValue{
						"Id": &types.AttributeValueMemberS{Value: id},
						"x":  &types.AttributeValueMemberN{Value: fmt.Sprint(x)},
						"y":  &types.AttributeValueMemberN{Value: fmt.Sprint(y)},
						"z":  &types.AttributeValueMemberN{Value: fmt.Sprint(z)},
					},
					TableName: aws.String("yofio"),
				},
			},
			{
				Update: &types.Update{
					ExpressionAttributeNames: map[string]string{
						"#aexcan": "AExCan",
						"#aeximp": "AExImp",
					},
					ExpressionAttributeValues: map[string]types.AttributeValue{
						":investment": &types.AttributeValueMemberN{Value: fmt.Sprint(req.Investment)},
						":uno":        &types.AttributeValueMemberN{Value: "1"},
					},
					Key: map[string]types.AttributeValue{
						"Id": &types.AttributeValueMemberS{Value: "*"},
					},
					TableName:        aws.String("yofio"),
					UpdateExpression: aws.String("SET #aexcan = #aexcan + :uno, #aeximp = #aeximp + :investment"),
				},
			},
		},
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf(`{"id": "%s", "credit_type_300": %d, "credit_type_500": %d, "credit_type_700": %d}`, id, x, y, z),
	}, nil
}

func main() {
	lambda.Start(insert)
}

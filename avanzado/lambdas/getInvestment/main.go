package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func findOne(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]
	if id == "*" || id == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "petición incorrecta"}`,
		}, nil
	}
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		}, nil
	}
	svc := dynamodb.NewFromConfig(cfg)
	res, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("yofio"),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{
				Value: id,
			},
		},
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		}, nil
	}
	if len(res.Item) == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       `{"error": "id ` + id + ` no existente"}`,
		}, nil
	}
	if res.Item["ANoExImp"] != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "id ` + id + ` `+ res.Item["ANoExImp"].(*types.AttributeValueMemberN).Value + ` no es distribuible"}`,
		}, nil
	} else {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body: `{"id": "` + id + `"` +
				`, "credit_type_300": ` + res.Item["x"].(*types.AttributeValueMemberN).Value +
				`, "credit_type_500": ` + res.Item["y"].(*types.AttributeValueMemberN).Value +
				`, "credit_type_700": ` + res.Item["z"].(*types.AttributeValueMemberN).Value + `}`,
		}, nil
	}
}

func main() {
	lambda.Start(findOne)
}

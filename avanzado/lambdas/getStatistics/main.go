package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Estadistica struct {
	AExCan,
	AExImp,
	ANoExCan,
	ANoExImp int32
}

func estadisticas(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
				Value: "*",
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
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error": "error recuperando estadísticas"}`,
		}, nil
	}
	estad := Estadistica{}
	err = attributevalue.UnmarshalMap(res.Item, &estad)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		}, nil
	}
	promexito := 0.0
	if estad.AExCan != 0 {
		promexito = float64(estad.AExImp) / float64(estad.AExCan)
	}
	promnoexito := 0.0
	if estad.ANoExCan != 0 {
		promnoexito = float64(estad.ANoExImp) / float64(estad.ANoExCan)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body: fmt.Sprintf(`{"asignaciones_realizadas": %d, "asignaciones_exitosas": %d, "asignaciones_no_exitosas": %d, `+
			`"promedio_inversión_exitosa": %.2f, "promedio_inversión_no_exitosa": %.2f}`,
			estad.AExCan+estad.ANoExCan, estad.AExCan, estad.ANoExCan, promexito, promnoexito),
	}, nil
}

func main() {
	lambda.Start(estadisticas)
}

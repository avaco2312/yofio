package initdb

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func CreateConfig(ctx context.Context) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return aws.Config{}, fmt.Errorf("error: Cargando configuracion AWS %v", err)
	}
	return cfg, nil
}

func CreateTable(ctx context.Context, cfg aws.Config) error {
	svc := dynamodb.NewFromConfig(cfg)
	_, err := svc.CreateTable(ctx,
		&dynamodb.CreateTableInput{
			TableName:   aws.String("yofio"), // Nombre de la tabla
			BillingMode: types.BillingModePayPerRequest,
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("Id"),
					KeyType:       types.KeyTypeHash, // Partition key
				},
			},
			AttributeDefinitions: []types.AttributeDefinition{
				{
					AttributeName: aws.String("Id"),
					AttributeType: types.ScalarAttributeTypeS,
				},
			},
		},
	)
	if err != nil {
		return fmt.Errorf("error: Creando tabla AWS %v", err)
	}
	waiter := dynamodb.NewTableExistsWaiter(svc)
	params := &dynamodb.DescribeTableInput{
		TableName: aws.String("yofio"),
	}
	maxWaitTime := 5 * time.Minute
	err = waiter.Wait(context.TODO(), params, maxWaitTime)
	if err != nil {
		return err
	}
	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("yofio"),
		Item: map[string]types.AttributeValue{
			"Id":       &types.AttributeValueMemberS{Value: "*"},
			"AExCan":   &types.AttributeValueMemberN{Value: "0"},
			"AExImp":   &types.AttributeValueMemberN{Value: "0"},
			"ANoExCan": &types.AttributeValueMemberN{Value: "0"},
			"ANoExImp": &types.AttributeValueMemberN{Value: "0"},
		},
	})
	return err
}

func DeleteTable(ctx context.Context, cfg aws.Config) error {
	svc := dynamodb.NewFromConfig(cfg)
	_, err := svc.DeleteTable(ctx,
		&dynamodb.DeleteTableInput{
			TableName: aws.String("yofio"),
		},
	)
	if err != nil {
		return err
	}	
	waiter := dynamodb.NewTableNotExistsWaiter(svc)
	params := &dynamodb.DescribeTableInput{
		TableName: aws.String("yofio"),
	}
	maxWaitTime := 5 * time.Minute
	err = waiter.Wait(context.TODO(), params, maxWaitTime)
	if err != nil {
		return err
	}
	return nil	
}

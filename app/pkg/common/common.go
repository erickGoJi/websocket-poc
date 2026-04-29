package common

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type HandlerFunc map[string]handlerFunc

type handlerFunc func(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error)

const (
	HandlerNameEnv   = "HANDLER_NAME"
	undefinedHandler = "undefined handler"
)

func RetrieveHandlerFunc(handlers map[string]handlerFunc) (handlerFunc, error) {
	handlerName := os.Getenv(HandlerNameEnv)

	if handler, ok := handlers[handlerName]; ok {
		return handler, nil
	}

	return nil, errors.New(undefinedHandler)
}

func RetrieveTableName() (string, error) {
	tableName := os.Getenv("TABLE_NAME")
	if tableName == "" {
		return "", errors.New("TABLE_NAME environment variable is not set")
	}
	return tableName, nil
}

func createAWSConfig() (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func InitializeDynamoDBClient(tableName *string) (*dynamodb.Client, error) {
	cfg, err := createAWSConfig()
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(*cfg)

	return client, nil
}

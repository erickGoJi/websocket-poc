package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/erickGoJi/websocket-lambda/pkg/common"
)

var TableName string
var DynamoDBClient *dynamodb.Client

var handlers = common.HandlerFunc{
	"ConnectHandler":    ConnectHandler,
	"DisconnectHandler": DisconnectHandler,
	"MessageHandler":    MessageHandler,
}

func initialize() {
	var err error
	TableName, err = common.RetrieveTableName()
	if err != nil {
		panic(err)
	}
	DynamoDBClient, err = common.InitializeDynamoDBClient(&TableName)
	if err != nil {
		panic(err)
	}
}

func main() {
	initialize()
	handler, err := common.RetrieveHandlerFunc(handlers)
	if err != nil {
		panic(err)
	}

	lambda.Start(handler)
}

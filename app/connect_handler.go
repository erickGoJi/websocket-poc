package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func ConnectHandler(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := DynamoDBClient.PutItem(
		ctx,
		&dynamodb.PutItemInput{
			TableName: &TableName,
			Item: map[string]types.AttributeValue{
				"connectionId": &types.AttributeValueMemberS{Value: request.RequestContext.ConnectionID},
				"createdAt":    &types.AttributeValueMemberS{Value: request.RequestContext.RequestTime},
			},
			ConditionExpression:    aws.String("attribute_not_exists(connectionId)"),
			ReturnValues:           types.ReturnValueNone,
			ReturnConsumedCapacity: types.ReturnConsumedCapacityTotal,
		},
	)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to store connection",
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Connected",
	}, nil
}

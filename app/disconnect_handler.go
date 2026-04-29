package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func DisconnectHandler(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := DynamoDBClient.DeleteItem(
		ctx,
		&dynamodb.DeleteItemInput{
			TableName: &TableName,
			Key: map[string]types.AttributeValue{
				"connectionId": &types.AttributeValueMemberS{Value: request.RequestContext.ConnectionID},
			},
			ConditionExpression: aws.String("attribute_exists(connectionId)"),
		},
	)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to delete connection",
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Disconnected",
	}, nil
}

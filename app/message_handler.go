package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

func MessageHandler(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	domain := request.RequestContext.DomainName
	stage := request.RequestContext.Stage
	connectionId := request.RequestContext.ConnectionID

	callbackURL := fmt.Sprintf("https://%s/%s", domain, stage)

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	apiClient := apigatewaymanagementapi.NewFromConfig(cfg, func(o *apigatewaymanagementapi.Options) {
		// ¡Paso CRÍTICO! Sobrescribimos el Endpoint con la URL de nuestro WebSocket
		o.BaseEndpoint = aws.String(callbackURL)
	})

	messageToSend := []byte("Hola! Recibí tu mensaje correctamente.")

	_, err = apiClient.PostToConnection(ctx, &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(connectionId),
		Data:         messageToSend,
	})
	if err != nil {
		fmt.Printf("Error enviando mensaje: %s\n", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Message received",
	}, nil
}

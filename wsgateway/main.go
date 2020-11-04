package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

// Handler ..
func Handler(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (*events.APIGatewayProxyResponse, error) {
	route := req.RequestContext.RouteKey
	switch route {
	case "$connect":
		fmt.Println("$connect was triggered")
	case "$disconnect":
		fmt.Println("$disconnect was triggered")
	case "$default":
		fmt.Println("$default was triggered")
	case "public/ping":
		region := os.Getenv("AWS_REGION")
		endpointURL := fmt.Sprintf("https://%s.execute-api.%s.amazonaws.com/dev", req.RequestContext.APIID, region)
		cfg := aws.NewConfig().WithEndpoint(endpointURL)
		apigwClient := apigatewaymanagementapi.New(session.New(), cfg)
		_, err := apigwClient.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: &req.RequestContext.ConnectionID,
			Data:         []byte(`{"body": "pong"}`),
		})
		if err != nil {
			return nil, err
		}
	}
	return &events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(Handler)
}

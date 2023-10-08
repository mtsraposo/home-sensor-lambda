package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request events.APIGatewayProxyRequest

func HandleRequest(ctx context.Context, request Request) (events.APIGatewayProxyResponse, error) {
	timestamp := request.Headers["timestamp"]
	fmt.Printf("Received timestamp: %s\n", timestamp)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Logged timestamp: %s", timestamp),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

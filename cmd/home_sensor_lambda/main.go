package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"home_sensor_lambda/internal/http_client"
	"home_sensor_lambda/pkg/config"
	"log"
)

type Request events.APIGatewayProxyRequest

func Process(request Request, httpClient http_client.Client) (int, string, error) {
	timestamp := request.Headers["X-Timestamp"]
	fmt.Printf("Received timestamp: %s\n", timestamp)
	appConfig, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	body, err := http_client.Get(appConfig.GithubUrl, httpClient)
	if err != nil {
		err := fmt.Errorf("request error: %v", err)
		fmt.Println(err)
		return 500, "", err
	}
	return 200, body, nil
}
func HandleRequest(_ctx context.Context, request Request) (events.APIGatewayProxyResponse, error) {
	httpClient := &http_client.HttpClient{}
	statusCode, body, _ := Process(request, httpClient)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"home_sensor_lambda/internal/http_client"
)

type Request events.APIGatewayProxyRequest

const url = "https://camo.githubusercontent.com/bc3509f02e6bcae7fd460fb19d6fae09ec3714c1b80339b99844335086572c4c/68747470733a2f2f6b6f6d617265762e636f6d2f67687076632f3f757365726e616d653d6d74737261706f736f"

func Process(request Request, httpClient http_client.Client) (int, string, error) {
	timestamp := request.Headers["X-Timestamp"]
	fmt.Printf("Received timestamp: %s\n", timestamp)
	body, err := http_client.Get(url, httpClient)
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

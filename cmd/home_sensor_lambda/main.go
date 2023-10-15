package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"home_sensor_lambda/internal/sns_client"
	"home_sensor_lambda/pkg/config"
	"log"
)

type Request events.APIGatewayProxyRequest

func Process(request Request, snsManager sns_client.SnsManager) (int, string, error) {
	timestamp := request.Headers["X-Timestamp"]
	fmt.Printf("Received timestamp: %s\n", timestamp)
	appConfig, configLoadError := config.Load()
	if configLoadError != nil {
		log.Fatalf("Failed to load configuration: %v", configLoadError)
	}
	publishPayload := sns_client.SnsPublishPayload{
		AwsRegion: appConfig.AwsRegion,
		Message:   timestamp,
		TopicArn:  appConfig.SnsTopicArn,
	}
	publishError := sns_client.Publish(snsManager, publishPayload)
	if publishError != nil {
		fmt.Println(fmt.Errorf("request error: %v", publishError))
		return 500, "", publishError
	}
	return 200, "OK", nil
}
func HandleRequest(_ctx context.Context, request Request) (events.APIGatewayProxyResponse, error) {
	snsManager := &sns_client.SnsBasicManager{}
	statusCode, body, _ := Process(request, snsManager)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

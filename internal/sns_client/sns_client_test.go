package sns_client_test

import (
	"home_sensor_lambda/internal/sns_client"
	"home_sensor_lambda/test/mocks"
	"testing"
)

func TestPublishSuccess(t *testing.T) {
	mockSnsPublisher := mocks.SnsManagerSuccess()
	publishPayload := sns_client.SnsPublishPayload{
		AwsRegion: "springfield",
		Message:   "hello world",
		TopicArn:  "arn:aws:iam::000000000000:role/presence_handler",
	}
	err := sns_client.Publish(mockSnsPublisher, publishPayload)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

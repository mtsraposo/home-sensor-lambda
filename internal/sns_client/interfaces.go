package sns_client

import (
	"github.com/aws/aws-sdk-go/service/sns"
)

type SnsPublishPayload struct {
	AwsRegion string
	Message   string
	TopicArn  string
}

type SnsPublisher interface {
	Publish(payload *sns.PublishInput) (*sns.PublishOutput, error)
}

type SnsManager interface {
	NewSession(awsRegion string) SnsPublisher
}

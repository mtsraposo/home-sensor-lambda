package sns_client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
)

type SnsBasicPublisher struct {
	publisher *sns.SNS
}

func (m *SnsBasicPublisher) Publish(payload *sns.PublishInput) (*sns.PublishOutput, error) {
	return m.publisher.Publish(payload)
}

type SnsBasicManager struct{}

func (c *SnsBasicManager) NewSession(awsRegion string) SnsPublisher {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	}))

	return &SnsBasicPublisher{
		publisher: sns.New(sess),
	}
}

func Publish(snsManager SnsManager, payload SnsPublishPayload) error {
	snsPublisher := snsManager.NewSession(payload.AwsRegion)
	result, err := snsPublisher.Publish(&sns.PublishInput{
		Message:  aws.String(payload.Message),
		TopicArn: aws.String(payload.TopicArn),
	})

	if err != nil {
		log.Println("Error publishing message:", err)
		return err
	}

	log.Println("Message published:", result.MessageId)
	return nil
}

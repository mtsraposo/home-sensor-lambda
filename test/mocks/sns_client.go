package mocks

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"home_sensor_lambda/internal/sns_client"
)

type MockSnsPublisher struct {
	MockPublish func(_payload *sns.PublishInput) (*sns.PublishOutput, error)
}

func (p *MockSnsPublisher) Publish(payload *sns.PublishInput) (*sns.PublishOutput, error) {
	return p.MockPublish(payload)
}

type MockSnsManager struct {
	MockNewSession func(awsRegion string) sns_client.SnsPublisher
}

func (m *MockSnsManager) NewSession(awsRegion string) sns_client.SnsPublisher {
	return m.MockNewSession(awsRegion)
}
func SnsManagerSuccess() *MockSnsManager {
	return &MockSnsManager{
		MockNewSession: func(awsRegion string) sns_client.SnsPublisher {
			return &MockSnsPublisher{
				MockPublish: func(_payload *sns.PublishInput) (*sns.PublishOutput, error) {
					return &sns.PublishOutput{}, nil
				},
			}
		},
	}
}

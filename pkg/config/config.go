package config

import (
	"errors"
	"os"
)

type Config struct {
	AwsRegion   string
	GithubUrl   string
	SnsTopicArn string
}

func Load() (*Config, error) {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev"
	}

	switch env {
	case "dev":
		return DevConfig()
	case "test":
		return TestConfig()
	case "prod":
		return ProdConfig()
	default:
		return nil, errors.New("invalid GO_ENV value")
	}
}

func DevConfig() (*Config, error) {
	awsRegion := os.Getenv("AWS_REGION")
	githubUrl := os.Getenv("GITHUB_URL")
	snsTopicArn := os.Getenv("SNS_TOPIC_ARN")
	return &Config{AwsRegion: awsRegion, GithubUrl: githubUrl, SnsTopicArn: snsTopicArn}, nil
}

func TestConfig() (*Config, error) {
	awsRegion := os.Getenv("AWS_REGION")
	githubUrl := os.Getenv("GITHUB_URL")
	snsTopicArn := os.Getenv("SNS_TOPIC_ARN")
	return &Config{AwsRegion: awsRegion, GithubUrl: githubUrl, SnsTopicArn: snsTopicArn}, nil
}

func ProdConfig() (*Config, error) {
	awsRegion := os.Getenv("AWS_REGION")
	githubUrl := os.Getenv("GITHUB_URL")
	snsTopicArn := os.Getenv("SNS_TOPIC_ARN")
	return &Config{AwsRegion: awsRegion, GithubUrl: githubUrl, SnsTopicArn: snsTopicArn}, nil
}

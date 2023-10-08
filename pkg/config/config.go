package config

import (
	"errors"
	"os"
)

type Config struct {
	GithubUrl string
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
	githubUrl := os.Getenv("GITHUB_URL")
	return &Config{GithubUrl: githubUrl}, nil
}

func TestConfig() (*Config, error) {
	githubUrl := os.Getenv("GITHUB_URL")
	return &Config{GithubUrl: githubUrl}, nil
}

func ProdConfig() (*Config, error) {
	githubUrl := os.Getenv("GITHUB_URL")
	return &Config{GithubUrl: githubUrl}, nil
}

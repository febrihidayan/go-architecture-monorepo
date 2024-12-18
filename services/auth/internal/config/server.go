package config

import (
	"os"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/config"
)

type AuthConfig struct {
	HttpPort     string
	RpcPort      string
	JwtTokenJti  string
	JwtExpired   int
	AppURL       string
	AppSecretKey string
	Timeout      time.Duration
	GrpcClient   GrpcClient
	RabbitMQ     RabbitMQConfig
}

type GrpcClient struct {
	User         string
	Notification string
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Exchange string
}

func Auth() *AuthConfig {
	return &AuthConfig{
		HttpPort:     os.Getenv("HTTP_PORT"),
		RpcPort:      os.Getenv("RPC_PORT"),
		JwtTokenJti:  os.Getenv("JWT_TOKEN_JTI"),
		AppURL:       os.Getenv("APP_URL"),
		AppSecretKey: os.Getenv("APP_SECRET_KEY"),
		JwtExpired:   config.ConvertInt("JWT_EXPIRED"),
		Timeout:      time.Duration(config.ConvertInt("APP_TIMEOUT")) * time.Second,
		GrpcClient: GrpcClient{
			User:         os.Getenv("RPC_USER"),
			Notification: os.Getenv("RPC_NOTIFICATION"),
		},
		RabbitMQ: RabbitMQConfig{
			Host:     os.Getenv("RABBITMQ_HOST"),
			Port:     os.Getenv("RABBITMQ_PORT"),
			User:     os.Getenv("RABBITMQ_USER"),
			Password: os.Getenv("RABBITMQ_PASSWORD"),
			Exchange: os.Getenv("RABBITMQ_EXCHANGE"),
		},
	}
}

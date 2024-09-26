package config

import (
	"os"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/config"
)

type UserConfig struct {
	HttpPort   string
	RpcPort    string
	Timeout    time.Duration
	GrpcClient GrpcClient
	RabbitMQ   RabbitMQConfig
}

type GrpcClient struct {
	Auth    string
	Storage string
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Exchange string
}

func User() *UserConfig {
	return &UserConfig{
		HttpPort: os.Getenv("HTTP_PORT"),
		RpcPort:  os.Getenv("RPC_PORT"),
		Timeout:  time.Duration(config.ConvertInt("APP_TIMEOUT")) * time.Second,
		GrpcClient: GrpcClient{
			Auth:    os.Getenv("RPC_AUTH"),
			Storage: os.Getenv("RPC_STORAGE"),
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

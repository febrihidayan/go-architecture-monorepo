package config

import (
	"os"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/config"
)

type NotificationConfig struct {
	HttpPort   string
	RpcPort    string
	Timeout    time.Duration
	GrpcClient GrpcClient
}

type GrpcClient struct {
	User string
}

func Notification() *NotificationConfig {
	return &NotificationConfig{
		HttpPort: os.Getenv("HTTP_PORT"),
		RpcPort:  os.Getenv("RPC_PORT"),
		Timeout:  time.Duration(config.ConvertInt("APP_TIMEOUT")) * time.Second,
		GrpcClient: GrpcClient{
			User: os.Getenv("RPC_USER"),
		},
	}
}

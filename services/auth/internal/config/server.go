package config

import (
	"os"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/config"
)

type AuthConfig struct {
	HttpPort   string
	RpcPort    string
	Timeout    time.Duration
	GrpcClient GrpcClient
}

type GrpcClient struct {
	User string
}

func Auth() *AuthConfig {
	return &AuthConfig{
		HttpPort: os.Getenv("HTTP_PORT"),
		RpcPort:  os.Getenv("RPC_PORT"),
		Timeout:  time.Duration(config.ConvertInt("APP_TIMEOUT")) * time.Second,
		GrpcClient: GrpcClient{
			User: os.Getenv("RPC_USER"),
		},
	}
}

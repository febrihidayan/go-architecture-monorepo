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
}

type GrpcClient struct {
	Auth string
}

func User() *UserConfig {
	return &UserConfig{
		HttpPort: os.Getenv("HTTP_PORT"),
		RpcPort:  os.Getenv("RPC_PORT"),
		Timeout:  time.Duration(config.ConvertInt("APP_TIMEOUT")) * time.Second,
		GrpcClient: GrpcClient{
			Auth: os.Getenv("RPC_AUTH"),
		},
	}
}

package config

import (
	"os"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/config"
)

type AuthConfig struct {
	HttpPort string
	RpcPort  string
	Timeout  time.Duration
}

func Auth() *AuthConfig {
	return &AuthConfig{
		HttpPort: os.Getenv("HTTP_PORT"),
		RpcPort:  os.Getenv("RPC_PORT"),
		Timeout:  time.Duration(config.ConvertInt("APP_TIMEOUT")) * time.Second,
	}
}

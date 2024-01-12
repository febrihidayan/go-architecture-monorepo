package config

import (
	"os"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/config"
)

type AuthConfig struct {
	HttpPort    string
	RpcPort     string
	JwtTokenJti string
	JwtExpired  int
	Timeout     time.Duration
	GrpcClient  GrpcClient
}

type GrpcClient struct {
	User         string
	Notification string
}

func Auth() *AuthConfig {
	return &AuthConfig{
		HttpPort:    os.Getenv("HTTP_PORT"),
		RpcPort:     os.Getenv("RPC_PORT"),
		JwtTokenJti: os.Getenv("JWT_TOKEN_JTI"),
		JwtExpired:  config.ConvertInt("JWT_EXPIRED"),
		Timeout:     time.Duration(config.ConvertInt("APP_TIMEOUT")) * time.Second,
		GrpcClient: GrpcClient{
			User:         os.Getenv("RPC_USER"),
			Notification: os.Getenv("RPC_NOTIFICATION"),
		},
	}
}

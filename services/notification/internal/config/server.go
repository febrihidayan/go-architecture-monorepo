package config

import (
	"os"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/config"
)

type NotificationConfig struct {
	HttpPort              string
	RpcPort               string
	Timeout               time.Duration
	GrpcClient            GrpcClient
	FirebaseGoogleService *FirebaseGoogle
}

type GrpcClient struct {
	User string
}

type FirebaseGoogle struct {
	Type                string
	ProjectId           string
	PrivateKeyId        string
	PrivateKey          string
	ClientEmail         string
	ClientId            string
	AuthUri             string
	TokenUri            string
	AuthProviderCertUrl string
	ClientCertUrl       string
	UniverseDomain      string
}

func Notification() *NotificationConfig {
	return &NotificationConfig{
		HttpPort: os.Getenv("HTTP_PORT"),
		RpcPort:  os.Getenv("RPC_PORT"),
		Timeout:  time.Duration(config.ConvertInt("APP_TIMEOUT")) * time.Second,
		GrpcClient: GrpcClient{
			User: os.Getenv("RPC_USER"),
		},
		FirebaseGoogleService: &FirebaseGoogle{
			Type:                os.Getenv("GOOGLE_FIREBASE_TYPE"),
			ProjectId:           os.Getenv("GOOGLE_FIREBASE_PROJECT_ID"),
			PrivateKeyId:        os.Getenv("GOOGLE_FIREBASE_PRIVATE_KEY_ID"),
			PrivateKey:          os.Getenv("GOOGLE_FIREBASE_PRIVATE_KEY"),
			ClientEmail:         os.Getenv("GOOGLE_FIREBASE_CLIENT_EMAIL"),
			ClientId:            os.Getenv("GOOGLE_FIREBASE_CLIENT_ID"),
			AuthUri:             os.Getenv("GOOGLE_FIREBASE_AUTH_URI"),
			TokenUri:            os.Getenv("GOOGLE_FIREBASE_TOKEN_URI"),
			AuthProviderCertUrl: os.Getenv("GOOGLE_FIREBASE_AUTH_PROVIDER_CERT_URL"),
			ClientCertUrl:       os.Getenv("GOOGLE_FIREBASE_CLIENT_CERT_URL"),
			UniverseDomain:      os.Getenv("GOOGLE_FIREBASE_UNIVERSE_DOMAIN"),
		},
	}
}

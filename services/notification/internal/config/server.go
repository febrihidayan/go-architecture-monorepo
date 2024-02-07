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
	Mailgun               *Mailgun
}

type GrpcClient struct {
	User string
}

type FirebaseGoogle struct {
	Path string
}

type Mailgun struct {
	MailFromDomain string
	MailFromName   string
	MailgunDomain  string
	MailgunSecret  string
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
			Path: os.Getenv("GOOGLE_FIREBASE_PATH"),
		},
		Mailgun: &Mailgun{
			MailFromDomain: os.Getenv("MAILGUN_FROM_DOMAIN"),
			MailFromName:   os.Getenv("MAILGUN_FROM_NAME"),
			MailgunDomain:  os.Getenv("MAILGUN_DOMAIN"),
			MailgunSecret:  os.Getenv("MAILGUN_SECRET"),
		},
	}
}

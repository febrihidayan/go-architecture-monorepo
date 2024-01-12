package services

import (
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
)

type ClientService struct {
	FirebaseGoogle *FirebaseGoogleService
	Mailgun        *MailgunService
}

func NewServiceHandler(cfg *config.NotificationConfig) *ClientService {
	// run firebase google service
	firebaseGoogleService, errFGService := NewFcmGoogleService(cfg.FirebaseGoogleService)
	if errFGService != nil {
		log.Fatalf("did not connect firebase service: %v", errFGService)
	}

	mailgun := NewMailgunClient(cfg.Mailgun)

	return &ClientService{
		FirebaseGoogle: firebaseGoogleService,
		Mailgun:        mailgun,
	}
}

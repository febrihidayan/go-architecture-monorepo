package services

import (
	"fmt"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/mailgun/mailgun-go"
)

// Ref: https://github.com/Golang-Coach/Lessons/blob/master/GoMailer/main.go
type MailgunService struct {
	Cfg *config.Mailgun
}

func NewMailgunClient(cfg *config.Mailgun) *MailgunService {
	return &MailgunService{
		Cfg: cfg,
	}
}

func (x *MailgunService) SendEmail(subject string, to []string, htmlMessage, textMessage string) (string, error) {
	// NewMailGun creates a new client instance.
	client := mailgun.NewMailgun(x.Cfg.MailgunDomain, x.Cfg.MailgunSecret)

	// Create message
	message := client.NewMessage(
		fmt.Sprintf("%s <%s>", x.Cfg.MailFromName, x.Cfg.MailFromDomain),
		subject,
		textMessage,
		to...,
	)

	// set html
	message.SetHtml(htmlMessage)

	// send message and get result
	_, id, err := client.Send(message)
	if err != nil {
		return "", err
	}

	return id, nil
}

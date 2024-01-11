package services

import (
	"fmt"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/mailgun/mailgun-go"
)

// Ref: https://github.com/Golang-Coach/Lessons/blob/master/GoMailer/main.go
type MailgunService struct {
	Client *mailgun.MailgunImpl
	Cfg    *config.Mailgun
}

func NewMailgunClient(cfg *config.Mailgun) *MailgunService {
	// NewMailGun creates a new client instance.
	client := mailgun.NewMailgun(cfg.MailgunDomain, cfg.MailgunSecret)

	return &MailgunService{
		Client: client,
		Cfg:    cfg,
	}
}

func (x *MailgunService) SendEmail(subject string, to []string, htmlMessage, textMessage string) (string, error) {
	// Create message
	message := x.Client.NewMessage(
		fmt.Sprintf("%s <%s>", x.Cfg.MailFromName, x.Cfg.MailFromDomain),
		subject,
		textMessage,
		to...,
	)

	// set html
	message.SetHtml(htmlMessage)

	// send message and get result
	_, id, err := x.Client.Send(message)
	if err != nil {
		return "", err
	}

	return id, nil
}

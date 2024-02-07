package services

import (
	"context"
	"path/filepath"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"google.golang.org/api/option"
)

type FirebaseGoogleService struct {
	App *firebase.App
}

func NewFcmGoogleService(cfg *config.FirebaseGoogle) (*FirebaseGoogleService, error) {
	path, err := filepath.Abs(cfg.Path)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsFile(path)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	return &FirebaseGoogleService{
		App: app,
	}, nil
}

func (x *FirebaseGoogleService) SendPushMessage(osName, deviceToken, title, body string) error {
	ctx := context.Background()
	client, err := x.App.Messaging(ctx)
	if err != nil {
		return err
	}

	var message *messaging.Message

	if strings.Contains(strings.ToUpper(osName), entities.DeviceOsAndroid) {
		message = &messaging.Message{
			Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},
			Token: deviceToken,
			Android: &messaging.AndroidConfig{
				Priority: "high",
			},
		}
	} else if strings.Contains(strings.ToUpper(osName), entities.DeviceOsIOS) {
		message = &messaging.Message{
			Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},
			Token: deviceToken,
			APNS: &messaging.APNSConfig{
				Payload: &messaging.APNSPayload{
					Aps: &messaging.Aps{
						ContentAvailable: true,
					},
				},
			},
		}
	} else {
		message = &messaging.Message{
			Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},
			Token: deviceToken,
		}
	}

	_, err = client.Send(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

func (x *FirebaseGoogleService) SendPushNotification(osName, deviceToken, title, body string, data map[string]string) error {
	ctx := context.Background()
	client, err := x.App.Messaging(ctx)
	if err != nil {
		return err
	}

	var message *messaging.Message

	if strings.Contains(strings.ToUpper(osName), entities.DeviceOsAndroid) {
		message = &messaging.Message{
			Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},
			Token: deviceToken,
			Android: &messaging.AndroidConfig{
				Priority: "high",
				Notification: &messaging.AndroidNotification{
					Icon: "ic_notif",
				},
			},
			Data: data,
		}
	} else if strings.Contains(strings.ToUpper(osName), entities.DeviceOsIOS) {
		message = &messaging.Message{
			Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},
			Token: deviceToken,
			APNS: &messaging.APNSConfig{
				Payload: &messaging.APNSPayload{
					Aps: &messaging.Aps{
						ContentAvailable: true,
						Alert: &messaging.ApsAlert{
							Title: title,
							Body:  body,
						},
					},
				},
			},
			Data: data,
		}
	} else {
		message = &messaging.Message{
			Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},
			Token: deviceToken,
			Data:  data,
		}
	}

	_, err = client.Send(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

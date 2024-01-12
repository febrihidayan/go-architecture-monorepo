package mappers

import (
	notificationPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/notification"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

func ToProtoNotificationSendParams(x entities.NotificationSends) *notificationPb.SendParams {
	return &notificationPb.SendParams{
		UserId:    x.UserId,
		Type:      x.TemplateName,
		Data:      x.Data,
		Services:  x.Services,
		PathEmail: x.PathEmail,
	}
}

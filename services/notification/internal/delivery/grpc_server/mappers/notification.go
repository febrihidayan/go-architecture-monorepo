package mappers

import (
	notificationPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/notification"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

func ToDomainNotificationParamsDto(x *notificationPb.SendParams) entities.NotificationSends {
	return entities.NotificationSends{
		UserId:       x.GetUserId(),
		TemplateName: x.GetType(),
		Data:         x.GetData(),
		Services:     x.GetServices(),
	}
}

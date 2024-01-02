package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/mongo/models"
)

func ToModelNotification(x *entities.Notification) *models.Notification {
	return &models.Notification{
		ID:        x.ID.String(),
		UserId:    x.UserId,
		Type:      x.Type,
		Data:      x.Data,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func ToDomainNotification(x *models.Notification) *entities.Notification {
	id, _ := common.StringToID(x.ID)
	return &entities.Notification{
		ID:        id,
		UserId:    x.UserId,
		Type:      x.Type,
		Data:      x.Data,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func ToListDomainNotification(models []*models.Notification) []*entities.Notification {
	data := make([]*entities.Notification, 0)
	for _, item := range models {
		data = append(data, ToDomainNotification(item))
	}
	return data
}

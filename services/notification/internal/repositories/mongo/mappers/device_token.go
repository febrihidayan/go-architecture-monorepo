package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/mongo/models"
)

func ToModelDeviceToken(x *entities.DeviceToken) *models.DeviceToken {
	return &models.DeviceToken{
		ID:        x.ID.String(),
		UserId:    x.UserId,
		Token:     x.Token,
		OsName:    x.OsName,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func ToDomainDeviceToken(x *models.DeviceToken) *entities.DeviceToken {
	id, _ := common.StringToID(x.ID)
	return &entities.DeviceToken{
		ID:        id,
		UserId:    x.UserId,
		Token:     x.Token,
		OsName:    x.OsName,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func ToListDomainDeviceToken(models []*models.DeviceToken) []*entities.DeviceToken {
	data := make([]*entities.DeviceToken, 0)
	for _, item := range models {
		data = append(data, ToDomainDeviceToken(item))
	}
	return data
}

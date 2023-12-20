package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/models"
)

func ToModelPermissionUser(x *entities.PermissionUser) *models.PermissionUser {
	return &models.PermissionUser{
		PermissionId: x.PermissionId,
		UserId:       x.UserId,
	}
}

func ToListModelPermissionUser(items []*entities.PermissionUser) []interface{} {
	data := make([]interface{}, 0)
	for _, item := range items {
		data = append(data, ToModelPermissionUser(item))
	}
	return data
}

func ToDomainPermissionUser(x *models.PermissionUser) *entities.PermissionUser {
	return &entities.PermissionUser{
		PermissionId: x.PermissionId,
		UserId:       x.UserId,
	}
}

func ToListDomainPermissionUser(models []*models.PermissionUser) []*entities.PermissionUser {
	data := make([]*entities.PermissionUser, 0)
	for _, item := range models {
		data = append(data, ToDomainPermissionUser(item))
	}
	return data
}

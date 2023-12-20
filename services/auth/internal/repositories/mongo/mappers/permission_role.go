package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/models"
)

func ToModelPermissionRole(x *entities.PermissionRole) *models.PermissionRole {
	return &models.PermissionRole{
		PermissionId: x.PermissionId,
		RoleId:       x.RoleId,
	}
}

func ToListModelPermissionRole(items []*entities.PermissionRole) []interface{} {
	data := make([]interface{}, 0)
	for _, item := range items {
		data = append(data, ToModelPermissionRole(item))
	}
	return data
}

func ToDomainPermissionRole(x *models.PermissionRole) *entities.PermissionRole {
	return &entities.PermissionRole{
		PermissionId: x.PermissionId,
		RoleId:       x.RoleId,
	}
}

func ToListDomainPermissionRole(models []*models.PermissionRole) []*entities.PermissionRole {
	data := make([]*entities.PermissionRole, 0)
	for _, item := range models {
		data = append(data, ToDomainPermissionRole(item))
	}
	return data
}

package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/models"
)

func ToModelRoleUser(x *entities.RoleUser) *models.RoleUser {
	return &models.RoleUser{
		UserId: x.UserId,
		RoleId: x.RoleId,
	}
}

func ToListModelRoleUser(items []*entities.RoleUser) []interface{} {
	data := make([]interface{}, 0)
	for _, item := range items {
		data = append(data, ToModelRoleUser(item))
	}
	return data
}

func ToDomainRoleUser(x *models.RoleUser) *entities.RoleUser {
	return &entities.RoleUser{
		UserId: x.UserId,
		RoleId: x.RoleId,
	}
}

func ToListDomainRoleUser(models []*models.RoleUser) []*entities.RoleUser {
	data := make([]*entities.RoleUser, 0)
	for _, item := range models {
		data = append(data, ToDomainRoleUser(item))
	}
	return data
}

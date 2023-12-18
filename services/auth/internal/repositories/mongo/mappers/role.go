package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/models"
)

func ToModelRole(x *entities.Role) *models.Role {
	return &models.Role{
		ID:          x.ID,
		Name:        x.Name,
		DisplayName: x.DisplayName,
		Description: x.Description,
		CreatedAt:   x.CreatedAt,
		UpdatedAt:   x.UpdatedAt,
	}
}

func ToDomainRole(x *models.Role) *entities.Role {
	return &entities.Role{
		ID:          x.ID,
		Name:        x.Name,
		DisplayName: x.DisplayName,
		Description: x.Description,
		CreatedAt:   x.CreatedAt,
		UpdatedAt:   x.UpdatedAt,
	}
}

func ToListDomainRole(models []*models.Role) []*entities.Role {
	data := make([]*entities.Role, 0)
	for _, item := range models {
		data = append(data, ToDomainRole(item))
	}
	return data
}

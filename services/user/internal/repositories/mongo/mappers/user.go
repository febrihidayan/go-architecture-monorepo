package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/mongo/models"
)

func ToModelUser(x *entities.User) *models.User {
	return &models.User{
		ID:        x.ID,
		Name:      x.Name,
		Email:     x.Email,
		Role:      x.Role,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func ToDomainUser(x *models.User) *entities.User {
	return &entities.User{
		ID:        x.ID,
		Name:      x.Name,
		Email:     x.Email,
		Role:      x.Role,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func ToListDomainUser(models []*models.User) []*entities.User {
	data := make([]*entities.User, 0)
	for _, item := range models {
		data = append(data, ToDomainUser(item))
	}
	return data
}

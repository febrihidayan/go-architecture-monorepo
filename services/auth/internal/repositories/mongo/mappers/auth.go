package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/models"
)

func ToModelAuth(x *entities.Auth) *models.Auth {
	return &models.Auth{
		ID:        x.ID,
		UserId:    x.UserId,
		Email:     x.Email,
		Password:  x.Password,
		Role:      x.Role,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func ToDomainAuth(x *models.Auth) *entities.Auth {
	return &entities.Auth{
		ID:        x.ID,
		UserId:    x.UserId,
		Email:     x.Email,
		Password:  x.Password,
		Role:      x.Role,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func ToListDomainAuth(models []*models.Auth) []*entities.Auth {
	data := make([]*entities.Auth, 0)
	for _, item := range models {
		data = append(data, ToDomainAuth(item))
	}
	return data
}

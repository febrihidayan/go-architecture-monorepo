package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/models"
)

func ToModelPermission(x *entities.Permission) *models.Permission {
	return &models.Permission{
		ID:          x.ID.String(),
		Name:        x.Name,
		DisplayName: x.DisplayName,
		Description: x.Description,
		CreatedAt:   x.CreatedAt,
		UpdatedAt:   x.UpdatedAt,
	}
}

func ToDomainPermission(x *models.Permission) *entities.Permission {
	id, _ := common.StringToID(x.ID)
	return &entities.Permission{
		ID:          id,
		Name:        x.Name,
		DisplayName: x.DisplayName,
		Description: x.Description,
		CreatedAt:   x.CreatedAt,
		UpdatedAt:   x.UpdatedAt,
	}
}

func ToListDomainPermission(models []*models.Permission) []*entities.Permission {
	data := make([]*entities.Permission, 0)
	for _, item := range models {
		data = append(data, ToDomainPermission(item))
	}
	return data
}

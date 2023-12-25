package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/mongo/models"
)

func ToModelCloud(x *entities.Cloud) *models.Cloud {
	return &models.Cloud{
		ID:        x.ID.String(),
		Name:      x.Name,
		Url:       x.Url,
		CreatedBy: x.CreatedBy,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func ToDomainCloud(x *models.Cloud) *entities.Cloud {
	id, _ := common.StringToID(x.ID)
	return &entities.Cloud{
		ID:        id,
		Name:      x.Name,
		Url:       x.Url,
		CreatedBy: x.CreatedBy,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func ToListDomainCloud(models []*models.Cloud) []*entities.Cloud {
	data := make([]*entities.Cloud, 0)
	for _, item := range models {
		data = append(data, ToDomainCloud(item))
	}
	return data
}

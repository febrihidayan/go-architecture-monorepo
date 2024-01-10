package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/mongo/models"
)

func ToModelTemplate(x *entities.Template) *models.Template {
	return &models.Template{
		ID:        x.ID.String(),
		Name:      x.Name,
		Type:      x.Type,
		Data:      x.Data,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func ToDomainTemplate(x *models.Template) *entities.Template {
	id, _ := common.StringToID(x.ID)
	return &entities.Template{
		ID:        id,
		Name:      x.Name,
		Type:      x.Type,
		Data:      x.Data,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func ToListDomainTemplate(models []*models.Template) []*entities.Template {
	data := make([]*entities.Template, 0)
	for _, item := range models {
		data = append(data, ToDomainTemplate(item))
	}
	return data
}

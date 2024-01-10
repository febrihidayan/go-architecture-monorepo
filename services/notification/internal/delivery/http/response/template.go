package response

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

type TemplateListResponse struct {
	ID        common.ID   `json:"id"`
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	Data      interface{} `json:"data"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func MapTemplateListResponse(x *entities.Template) TemplateListResponse {
	return TemplateListResponse{
		ID:        x.ID,
		Name:      x.Name,
		Type:      x.Type,
		Data:      x.GetData(),
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func MapTemplateListResponses(x []*entities.Template) []TemplateListResponse {
	var result []TemplateListResponse
	for _, v := range x {
		result = append(result, MapTemplateListResponse(v))
	}
	return result
}

package response

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
)

type CloudListResponse struct {
	ID        common.ID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapCloudListResponse(x *entities.Cloud) CloudListResponse {
	return CloudListResponse{
		ID:        x.ID,
		Name:      x.Name,
		Url:       x.Url,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func MapCloudListResponses(x []*entities.Cloud) []CloudListResponse {
	var result []CloudListResponse
	for _, v := range x {
		result = append(result, MapCloudListResponse(v))
	}
	return result
}

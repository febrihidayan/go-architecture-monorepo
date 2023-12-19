package response

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type PermissionListResponse struct {
	ID          common.ID `json:"id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAtAt time.Time `json:"updated_at"`
}

func MapPermissionListResponse(x *entities.Permission) PermissionListResponse {
	return PermissionListResponse{
		ID:          x.ID,
		Name:        x.Name,
		DisplayName: x.DisplayName,
		Description: x.Description,
		CreatedAt:   x.CreatedAt,
	}
}

func MapPermissionListResponses(x []*entities.Permission) []PermissionListResponse {
	var result []PermissionListResponse
	for _, v := range x {
		result = append(result, MapPermissionListResponse(v))
	}
	return result
}

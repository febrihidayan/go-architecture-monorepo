package response

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type RoleListResponse struct {
	ID          common.ID `json:"id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func MapRoleListResponse(x *entities.Role) RoleListResponse {
	return RoleListResponse{
		ID:          x.ID,
		Name:        x.Name,
		DisplayName: x.DisplayName,
		Description: x.Description,
		CreatedAt:   x.CreatedAt,
		UpdatedAt:   x.UpdatedAt,
	}
}

func MapRoleListResponses(x []*entities.Role) []RoleListResponse {
	var result []RoleListResponse
	for _, v := range x {
		result = append(result, MapRoleListResponse(v))
	}
	return result
}

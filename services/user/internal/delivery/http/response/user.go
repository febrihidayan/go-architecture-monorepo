package response

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
)

type UserListResponse struct {
	ID        common.ID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	LangCode  string    `json:"lang_code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapUserListResponse(x *entities.User) UserListResponse {
	return UserListResponse{
		ID:        x.ID,
		Name:      x.Name,
		Email:     x.Email,
		Avatar:    x.Avatar,
		LangCode:  x.LangCode,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

func MapUserListResponses(x []*entities.User) []UserListResponse {
	var result []UserListResponse
	for _, v := range x {
		result = append(result, MapUserListResponse(v))
	}
	return result
}

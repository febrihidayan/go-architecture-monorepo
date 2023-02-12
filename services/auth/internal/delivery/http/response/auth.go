package response

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type LoginResponse struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapLoginResponse(x *entities.Auth) LoginResponse {
	return LoginResponse{
		ID:        x.ID.String(),
		UserId:    x.UserId,
		Email:     x.Email,
		Role:      x.Role,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

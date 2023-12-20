package response

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type LoginToken struct {
	Sub   string   `json:"sub"`
	JTI   string   `json:"jti"`
	Roles []string `json:"roles,omitempty"`
	Exp   int      `json:"exp"`
}

type LoginResponse struct {
	AccessToken  LoginToken `json:"access_token"`
	RefreshToken LoginToken `json:"refresh_token"`
	Exp          int        `json:"exp"`
}

func MapLoginResponse(x *entities.AuthTokenMeta) LoginResponse {
	return LoginResponse{
		AccessToken: LoginToken{
			Sub:   x.AccessToken.Sub,
			JTI:   x.AccessToken.JTI,
			Roles: x.AccessToken.Roles,
			Exp:   x.AccessToken.Exp,
		},
		RefreshToken: LoginToken{
			Sub: x.AccessToken.Sub,
			JTI: x.AccessToken.JTI,
			Exp: x.AccessToken.Exp,
		},
		Exp: x.Exp,
	}
}

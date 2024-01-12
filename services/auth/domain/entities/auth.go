package entities

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"

	"github.com/hashicorp/go-multierror"
)

type Auth struct {
	ID        common.ID
	UserId    string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthDto struct {
	ID       *common.ID
	UserId   string
	Email    string
	Password string
}

type AuthMeta struct {
	Auth  *Auth
	Roles []string
}

type AuthToken struct {
	Sub   string
	JTI   string
	Roles []string
	Exp   int
}

type AuthTokenMeta struct {
	AccessToken  AuthToken
	RefreshToken AuthToken
	Exp          int
}

func NewAuth(x AuthDto, finds ...*Auth) *Auth {
	auth := Auth{
		ID:        common.NewID(),
		UserId:    x.UserId,
		Email:     x.Email,
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	if x.ID != nil {
		auth.ID = *x.ID
	}

	for _, item := range finds {
		if item != nil {
			auth.Password = item.Password
			auth.CreatedAt = item.CreatedAt
		}
	}

	// set password hash
	if x.Password != "" {
		auth.SetPasswordHash(x.Password)
	}

	return &auth
}

func NewAuthLogin(x *AuthMeta) *AuthTokenMeta {
	duration := config.Auth().JwtExpired
	tokenJTI := config.Auth().JwtTokenJti

	exp := time.Now().Add(time.Second * time.Duration(duration) * 60 * 24).Unix() // 1 day
	return &AuthTokenMeta{
		AccessToken: AuthToken{
			Sub:   x.Auth.UserId,
			JTI:   tokenJTI,
			Exp:   int(exp),
			Roles: x.Roles,
		},
		RefreshToken: AuthToken{
			Sub: x.Auth.UserId,
			JTI: tokenJTI,
			Exp: int(exp),
		},
		Exp: int(exp),
	}
}

func (x *Auth) Validate() (err *multierror.Error) {
	if x.UserId == "" {
		err = multierror.Append(err, lang.ErrUserIdRequired)
	}
	if x.Email == "" {
		err = multierror.Append(err, lang.ErrEmailRequired)
	}
	if x.Password == "" {
		err = multierror.Append(err, lang.ErrPasswordRequired)
	}

	return
}

func (x *Auth) SetPasswordHash(hashedPwd string) {
	pwd, _ := utils.HashPassword(hashedPwd)
	x.Password = pwd
}

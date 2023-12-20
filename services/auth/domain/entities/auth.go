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
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthDto struct {
	ID       *common.ID
	UserId   string
	Email    string
	Password string
	Role     string
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

func NewAuth(x AuthDto) *Auth {
	id := common.NewID()

	if x.ID != nil {
		id = *x.ID
	}

	return &Auth{
		ID:        id,
		UserId:    x.UserId,
		Email:     x.Email,
		Password:  x.Password,
		Role:      x.Role,
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}
}

func NewAuthLogin(x *Auth) *AuthTokenMeta {
	duration := config.Auth().JwtExpired
	tokenJTI := config.Auth().JwtTokenJti

	exp := time.Now().Add(time.Second * time.Duration(duration) * 60 * 24).Unix() // 1 day
	return &AuthTokenMeta{
		AccessToken: AuthToken{
			Sub:   x.UserId,
			JTI:   tokenJTI,
			Exp:   int(exp),
			Roles: []string{x.Role},
		},
		RefreshToken: AuthToken{
			Sub: x.UserId,
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
	if x.Role == "" {
		err = multierror.Append(err, lang.ErrRoleRequired)
	}

	return
}

func (x *Auth) SetPasswordHash(hashedPwd string) {
	pwd, _ := utils.HashPassword(hashedPwd)
	x.Password = pwd
}

package entities

import (
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"

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

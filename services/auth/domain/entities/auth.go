package entities

import (
	"errors"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
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
		err = multierror.Append(err, errors.New("UserId is required"))
	}
	if x.Email == "" {
		err = multierror.Append(err, errors.New("Email is required"))
	}
	if x.Password == "" {
		err = multierror.Append(err, errors.New("Password is required"))
	}
	if x.Role == "" {
		err = multierror.Append(err, errors.New("Role is required"))
	}

	return
}

func (x *Auth) SetPasswordHash(hashedPwd string) {
	pwd, _ := utils.HashPassword(hashedPwd)
	x.Password = pwd
}

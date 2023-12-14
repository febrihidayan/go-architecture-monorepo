package entities

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
)

type Register struct {
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
}

type RegisterDto struct {
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
}

func NewRegister(x RegisterDto) *Register {
	return &Register{
		Name:            x.Name,
		Email:           x.Email,
		Password:        x.Password,
		ConfirmPassword: x.ConfirmPassword,
	}
}

func (x *Register) Validate() (err *multierror.Error) {
	if x.Name == "" {
		err = multierror.Append(err, lang.ErrNameRequired)
	}
	if x.Email == "" {
		err = multierror.Append(err, lang.ErrEmailRequired)
	}
	if x.Password == "" {
		err = multierror.Append(err, lang.ErrPasswordRequired)
	}
	if x.ConfirmPassword == "" {
		err = multierror.Append(err, lang.ErrConfirmPasswordRequired)
	}
	if x.ConfirmPassword != x.Password {
		err = multierror.Append(err, lang.ErrConfitmPasswordNotSame)
	}

	return
}

package entities

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
)

type RoleUser struct {
	UserId string
	RoleId string
}

type RoleUserDto struct {
	UserId string
	RoleId string
}

func NewRoleUser(x RoleUserDto) *RoleUser {
	return &RoleUser{
		UserId: x.UserId,
		RoleId: x.RoleId,
	}
}

func (x *RoleUser) Validate() (err *multierror.Error) {
	if x.UserId == "" {
		err = multierror.Append(err, lang.Trans("filled", "UserId"))
	}
	if x.RoleId == "" {
		err = multierror.Append(err, lang.Trans("filled", "RoleId"))
	}

	return
}

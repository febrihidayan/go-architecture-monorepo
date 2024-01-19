package entities

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
)

type RoleUser struct {
	ID     common.ID
	UserId string
	RoleId string
}

type RoleUserDto struct {
	ID     *common.ID
	UserId string
	RoleId string
}

func NewRoleUser(x RoleUserDto) *RoleUser {
	roleUser := RoleUser{
		UserId: x.UserId,
		RoleId: x.RoleId,
	}

	if x.ID != nil {
		roleUser.ID = *x.ID
	}

	return &roleUser
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

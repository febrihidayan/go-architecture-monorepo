package entities

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
)

type PermissionUser struct {
	PermissionId string
	UserId       string
}

type PermissionUserDto struct {
	PermissionId string
	UserId       string
}

func NewPermissionUser(x PermissionUserDto) *PermissionUser {
	return &PermissionUser{
		PermissionId: x.PermissionId,
		UserId:       x.UserId,
	}
}

func (x *PermissionUser) Validate() (err *multierror.Error) {
	if x.PermissionId == "" {
		err = multierror.Append(err, lang.Trans("filled", "PermissionId"))
	}
	if x.UserId == "" {
		err = multierror.Append(err, lang.Trans("filled", "UserId"))
	}

	return
}

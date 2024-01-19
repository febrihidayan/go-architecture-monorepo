package entities

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
)

type PermissionUser struct {
	ID           common.ID
	PermissionId string
	UserId       string
}

type PermissionUserDto struct {
	ID           *common.ID
	PermissionId string
	UserId       string
}

func NewPermissionUser(x PermissionUserDto) *PermissionUser {
	permissionUser := PermissionUser{
		PermissionId: x.PermissionId,
		UserId:       x.UserId,
	}

	if x.ID != nil {
		permissionUser.ID = *x.ID
	}

	return &permissionUser
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

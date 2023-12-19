package entities

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
)

type PermissionRole struct {
	PermissionId string
	RoleId       string
}

type PermissionRoleDto struct {
	PermissionId string
	RoleId       string
}

func NewPermissionRole(x PermissionRoleDto) *PermissionRole {
	return &PermissionRole{
		PermissionId: x.PermissionId,
		RoleId:       x.RoleId,
	}
}

func (x *PermissionRole) Validate() (err *multierror.Error) {
	if x.PermissionId == "" {
		err = multierror.Append(err, lang.Trans("filled", "PermissionId"))
	}
	if x.RoleId == "" {
		err = multierror.Append(err, lang.Trans("filled", "RoleId"))
	}

	return
}

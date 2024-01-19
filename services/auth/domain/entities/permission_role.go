package entities

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
)

type PermissionRole struct {
	ID           common.ID
	PermissionId string
	RoleId       string
}

type PermissionRoleDto struct {
	ID           *common.ID
	PermissionId string
	RoleId       string
}

func NewPermissionRole(x PermissionRoleDto) *PermissionRole {
	permissionRole := PermissionRole{
		PermissionId: x.PermissionId,
		RoleId:       x.RoleId,
	}

	if x.ID != nil {
		permissionRole.ID = *x.ID
	}

	return &permissionRole
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

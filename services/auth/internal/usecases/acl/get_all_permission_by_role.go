package acl

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *aclInteractor) GetAllPermissionByRole(ctx context.Context, roleId string) ([]*entities.Permission, *exceptions.CustomError) {
	var (
		multilerr   *multierror.Error
		permissions []*entities.Permission
	)

	permissionRole, err := x.permissionRoleRepo.AllByRoleId(ctx, roleId)
	if err != nil {
		log.Println("GetAllPermissionByRole::error#1:", err)
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	for _, item := range permissionRole {
		permission, _ := x.permissionRepo.Find(ctx, item.PermissionId)
		if permission != nil {
			permissions = append(permissions, permission)
		}
	}

	log.Println("GetAllPermissionByRole::success#1:", "successfully")

	return permissions, nil
}

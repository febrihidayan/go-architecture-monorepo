package acl

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/exp/slices"
)

func (x *aclInteractor) UpdatePermissionByRole(ctx context.Context, payload entities.AclPermissionDto) *exceptions.CustomError {
	var multilerr *multierror.Error

	role, err := x.roleRepo.Find(ctx, payload.RoleId)
	if err != nil {
		log.Println("UpdatePermissionByRole::error#1:", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	allPermissions, err := x.permissionRoleRepo.AllByRoleId(ctx, role.ID.String())
	if err != nil {
		log.Println("UpdatePermissionByRole::error#2:", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	if len(payload.Permissions) > 0 {
		var deletes []string

		for _, item := range allPermissions {
			if !slices.Contains(payload.Permissions, item.PermissionId) {
				deletes = append(deletes, item.PermissionId)
			} else {
				payload.Permissions[slices.Index(payload.Permissions, item.PermissionId)] = payload.Permissions[len(payload.Permissions)-1]
				payload.Permissions = payload.Permissions[:len(payload.Permissions)-1]
			}
		}

		permissionRoles := make([]*entities.PermissionRole, 0)
		for _, item := range payload.Permissions {
			permissionRole := entities.NewPermissionRole(entities.PermissionRoleDto{
				PermissionId: item,
				RoleId:       role.ID.String(),
			})
			permissionRoles = append(permissionRoles, permissionRole)
		}

		if len(permissionRoles) > 0 {
			if err := x.permissionRoleRepo.CreateMany(ctx, permissionRoles); err != nil {
				log.Println("UpdatePermissionByRole::error#3:", err)
				multilerr = multierror.Append(multilerr, err)
				return &exceptions.CustomError{
					Status: exceptions.ERRREPOSITORY,
					Errors: multilerr,
				}
			}
		}

		if len(deletes) > 0 {
			if err := x.permissionRoleRepo.DeleteByPermissionIds(ctx, deletes); err != nil {
				log.Println("UpdatePermissionByRole::error#4:", err)
				multilerr = multierror.Append(multilerr, err)
				return &exceptions.CustomError{
					Status: exceptions.ERRREPOSITORY,
					Errors: multilerr,
				}
			}
		}

	} else if len(allPermissions) > 0 {
		if err := x.permissionRoleRepo.DeleteByRoleId(ctx, role.ID.String()); err != nil {
			log.Println("UpdatePermissionByRole::error#5:", err)
			multilerr = multierror.Append(multilerr, err)
			return &exceptions.CustomError{
				Status: exceptions.ERRREPOSITORY,
				Errors: multilerr,
			}
		}
	}

	log.Println("UpdatePermissionByRole::success#1:", "successfully")

	return nil
}

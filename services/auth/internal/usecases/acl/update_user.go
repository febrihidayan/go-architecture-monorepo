package acl

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/exp/slices"
)

func (x *aclInteractor) UpdateUser(ctx context.Context, payload entities.AclUserDto) *exceptions.CustomError {
	var multilerr *multierror.Error

	user, err := x.authRepo.FindByUserId(ctx, payload.UserId)
	if err != nil {
		log.Println("UpdateUser::error#1:", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	allRoles, err := x.roleUserRepo.AllByUserId(ctx, user.UserId)
	if err != nil {
		log.Println("UpdateUser::error#2:", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	allPermissions, err := x.permissionUserRepo.AllByUserId(ctx, user.UserId)
	if err != nil {
		log.Println("UpdateUser::error#3:", err)
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

		permissionUsers := make([]*entities.PermissionUser, 0)
		for _, item := range payload.Permissions {
			permissionUser := entities.NewPermissionUser(entities.PermissionUserDto{
				PermissionId: item,
				UserId:       user.UserId,
			})
			permissionUsers = append(permissionUsers, permissionUser)
		}

		if len(permissionUsers) > 0 {
			if err := x.permissionUserRepo.CreateMany(ctx, permissionUsers); err != nil {
				log.Println("UpdateUser::error#4:", err)
				multilerr = multierror.Append(multilerr, err)
				return &exceptions.CustomError{
					Status: exceptions.ERRREPOSITORY,
					Errors: multilerr,
				}
			}
		}

		if len(deletes) > 0 {
			if err := x.permissionUserRepo.DeleteByPermissionIds(ctx, deletes); err != nil {
				log.Println("UpdateUser::error#5:", err)
				multilerr = multierror.Append(multilerr, err)
				return &exceptions.CustomError{
					Status: exceptions.ERRREPOSITORY,
					Errors: multilerr,
				}
			}
		}

	} else if len(allPermissions) > 0 {
		if err := x.permissionUserRepo.DeleteByUserId(ctx, user.UserId); err != nil {
			log.Println("UpdateUser::error#6:", err)
			multilerr = multierror.Append(multilerr, err)
			return &exceptions.CustomError{
				Status: exceptions.ERRREPOSITORY,
				Errors: multilerr,
			}
		}
	}

	if len(payload.Roles) > 0 {
		var deletes []string

		for _, item := range allRoles {
			if !slices.Contains(payload.Roles, item.RoleId) {
				deletes = append(deletes, item.RoleId)
			} else {
				payload.Roles[slices.Index(payload.Roles, item.RoleId)] = payload.Roles[len(payload.Roles)-1]
				payload.Roles = payload.Roles[:len(payload.Roles)-1]
			}
		}

		roleUsers := make([]*entities.RoleUser, 0)
		for _, item := range payload.Roles {
			roleUser := entities.NewRoleUser(entities.RoleUserDto{
				RoleId: item,
				UserId: user.UserId,
			})

			roleUsers = append(roleUsers, roleUser)
		}

		if len(roleUsers) > 0 {
			if err := x.roleUserRepo.CreateMany(ctx, roleUsers); err != nil {
				log.Println("UpdateUser::error#7:", err)
				multilerr = multierror.Append(multilerr, err)
				return &exceptions.CustomError{
					Status: exceptions.ERRREPOSITORY,
					Errors: multilerr,
				}
			}
		}

		if len(deletes) > 0 {
			if err := x.roleUserRepo.DeleteByRoleIds(ctx, deletes); err != nil {
				log.Println("UpdateUser::error#8:", err)
				multilerr = multierror.Append(multilerr, err)
				return &exceptions.CustomError{
					Status: exceptions.ERRREPOSITORY,
					Errors: multilerr,
				}
			}
		}

	} else if len(allRoles) > 0 {
		if err := x.roleUserRepo.DeleteByUserId(ctx, user.UserId); err != nil {
			log.Println("UpdateUser::error#9:", err)
			multilerr = multierror.Append(multilerr, err)
			return &exceptions.CustomError{
				Status: exceptions.ERRREPOSITORY,
				Errors: multilerr,
			}
		}
	}

	log.Println("UpdateUser::success#1:", "saved successfully")

	return nil
}

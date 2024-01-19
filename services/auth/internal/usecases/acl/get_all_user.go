package acl

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *aclInteractor) GetAllUser(ctx context.Context, userId string) (*entities.AclMeta, *exceptions.CustomError) {
	var multilerr *multierror.Error

	user, err := x.authRepo.FindByUserId(ctx, userId)
	if err != nil {
		log.Println("GetAllUser::error#1:", err)
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	roles, err := x.roleRepo.AllByUserId(ctx, user.UserId)
	if err != nil {
		log.Println("GetAllUser::error#2:", err)
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	permissions, err := x.permissionRepo.AllByUserId(ctx, user.UserId)
	if err != nil {
		log.Println("GetAllUser::error#3:", err)
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	log.Println("GetAllUser::success#1:", "successfully")

	return &entities.AclMeta{
		Roles:       roles,
		Permissions: permissions,
	}, nil
}

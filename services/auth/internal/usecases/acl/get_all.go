package acl

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *aclInteractor) GetAll(ctx context.Context) (*entities.AclMeta, *exceptions.CustomError) {
	var multilerr *multierror.Error

	roles, err := x.roleRepo.All(ctx)
	if err != nil {
		log.Println("GetAll::error#1:", err)
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	permissions, err := x.permissionRepo.All(ctx)
	if err != nil {
		log.Println("GetAll::error#2:", err)
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	log.Println("GetAll::success#1:", "successfully")

	return &entities.AclMeta{
		Roles:       roles,
		Permissions: permissions,
	}, nil
}

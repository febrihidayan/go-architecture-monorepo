package acl

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *aclInteractor) GetAllPermission(ctx context.Context) ([]*entities.Permission, *exceptions.CustomError) {
	var multilerr *multierror.Error

	permissions, err := x.permissionRepo.All(ctx)
	if err != nil {
		log.Println("GetAllPermission::error#1:", err)
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	log.Println("GetAllPermission::success#1:", "successfully")

	return permissions, nil
}

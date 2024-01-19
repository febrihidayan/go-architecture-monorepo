package acl

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *aclInteractor) GetAllRole(ctx context.Context) ([]*entities.Role, *exceptions.CustomError) {
	var multilerr *multierror.Error

	roles, err := x.roleRepo.All(ctx)
	if err != nil {
		log.Println("GetAllRole::error#1:", err)
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	log.Println("GetAllRole::success#1:", "successfully")

	return roles, nil
}

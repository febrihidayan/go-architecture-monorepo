package permission

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *permissionInteractor) Find(ctx context.Context, id string) (*entities.Permission, *exceptions.CustomError) {
	var multilerr *multierror.Error

	find, err := x.permissionRepo.Find(ctx, id)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return find, nil
}

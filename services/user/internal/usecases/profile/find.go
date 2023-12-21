package profile

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"

	"github.com/hashicorp/go-multierror"
)

func (x *profileInteractor) Find(ctx context.Context, id string) (*entities.User, *exceptions.CustomError) {
	var multilerr *multierror.Error

	find, err := x.userRepo.Find(ctx, id)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return find, nil
}

package user

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (x *userInteractor) Delete(ctx context.Context, id string) *exceptions.CustomError {
	var multilerr *multierror.Error

	if err := x.userRepo.Delete(ctx, id); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	if err := x.rabbitmqRepo.AuthDelete(id); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	return nil
}

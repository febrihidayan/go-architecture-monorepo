package user

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"

	"github.com/hashicorp/go-multierror"
)

func (x *userInteractor) GetAll(ctx context.Context, params entities.UserQueryParams) (*entities.UserMeta, *exceptions.CustomError) {
	var multilerr *multierror.Error

	all, total, err := x.userRepo.GetAll(ctx, &params)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return &entities.UserMeta{
		Data:  all,
		Total: total,
	}, nil
}

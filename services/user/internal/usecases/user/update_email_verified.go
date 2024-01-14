package user

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"

	"github.com/hashicorp/go-multierror"
)

func (x *userInteractor) UpdateEmailVerified(ctx context.Context, payload entities.UserDto) *exceptions.CustomError {
	var multilerr *multierror.Error

	user, err := x.userRepo.Find(ctx, payload.ID.String())
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	user.SetEmailVerifiedAt(payload.EmailVerifiedAt)

	if err := x.userRepo.Update(ctx, user); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return nil
}

package user

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"

	"github.com/hashicorp/go-multierror"
)

func (x *userInteractor) Update(ctx context.Context, payload entities.UserDto) (*entities.User, *exceptions.CustomError) {
	var multilerr *multierror.Error

	find, err := x.userRepo.Find(ctx, payload.ID.String())
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	user := entities.NewUser(payload, find)
	if err := user.Validate(); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRDOMAIN,
			Errors: multilerr,
		}
	}

	auth := entities.Auth{
		UserId:   user.ID.String(),
		Email:    user.Email,
		Password: payload.Auth.Password,
	}

	if err := x.authGrpcRepo.CreateOrUpdate(ctx, &auth); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	if err := x.userRepo.Update(ctx, user); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	if user.Avatar != "" && find.Avatar != user.Avatar {
		if err := x.storageGrpcRepo.UpdateCloudApprove(ctx, []string{user.Avatar}); err != nil {
			multilerr = multierror.Append(multilerr, err)
			return nil, &exceptions.CustomError{
				Status: exceptions.ERRREPOSITORY,
				Errors: multilerr,
			}
		}
	}

	if find.Avatar != "" && find.Avatar != user.Avatar {
		if err := x.storageGrpcRepo.DeleteCloudApprove(ctx, []string{find.Avatar}); err != nil {
			multilerr = multierror.Append(multilerr, err)
			return nil, &exceptions.CustomError{
				Status: exceptions.ERRREPOSITORY,
				Errors: multilerr,
			}
		}
	}

	return user, nil
}

package user

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"

	"github.com/hashicorp/go-multierror"
)

func (x *userInteractor) Create(ctx context.Context, payload entities.UserDto) (*entities.User, *exceptions.CustomError) {
	var multilerr *multierror.Error

	find, _ := x.userRepo.FindByEmail(ctx, payload.Email)
	if find != nil {
		multilerr = multierror.Append(multilerr, lang.ErrEmailAlready)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	user := entities.NewUser(payload)
	user.DefaultLang()

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

	if err := x.userRepo.Create(ctx, user); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	if user.Avatar != "" {
		if err := x.storageGrpcRepo.UpdateCloudApprove(ctx, []string{user.Avatar}); err != nil {
			multilerr = multierror.Append(multilerr, err)
			return nil, &exceptions.CustomError{
				Status: exceptions.ERRREPOSITORY,
				Errors: multilerr,
			}
		}
	}

	return user, nil
}

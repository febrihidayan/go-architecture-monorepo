package auth

import (
	"context"
	"errors"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"

	"github.com/hashicorp/go-multierror"
)

func (x *authInteractor) Register(ctx context.Context, payload entities.RegisterDto) (*entities.Auth, *exceptions.CustomError) {
	var multilerr *multierror.Error

	log.Println("start check email already")
	find, _ := x.authRepo.FindByEmail(ctx, payload.Email)
	if find != nil {
		multilerr = multierror.Append(multilerr, errors.New("The email is already registered."))
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	log.Println("start create user")
	user, err := x.userRepo.CreateUser(ctx, entities.User{
		Name:  payload.Name,
		Email: payload.Email,
	})
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	log.Println("start auth dto")
	auth := entities.NewAuth(entities.AuthDto{
		UserId:   user.ID.String(),
		Email:    payload.Email,
		Password: payload.Password,
	})
	log.Println("start validation auth")
	if err := auth.Validate(); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRDOMAIN,
			Errors: multilerr,
		}
	}

	// set password
	auth.SetPasswordHash(payload.Password)

	log.Println("start create auth")
	if err := x.authRepo.Create(ctx, auth); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	log.Println("success create user and auth")

	return auth, nil
}

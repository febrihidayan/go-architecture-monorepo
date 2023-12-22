package auth

import (
	"context"
	"errors"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/middleware"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"

	"github.com/hashicorp/go-multierror"
)

func (x *authInteractor) CreateOrUpdate(ctx context.Context, payload entities.AuthDto) (*entities.Auth, *exceptions.CustomError) {
	var multilerr *multierror.Error

	find, _ := x.authRepo.FindByUserId(ctx, payload.UserId)

	if find != nil {
		// check email ready, if email not same
		if find.Email != payload.Email {
			findEmail, _ := x.authRepo.FindByEmail(ctx, payload.Email)
			if findEmail != nil {
				multilerr = multierror.Append(multilerr, errors.New("The email is already registered."))
				return nil, &exceptions.CustomError{
					Status: exceptions.ERRREPOSITORY,
					Errors: multilerr,
				}
			}
		}

		payload.ID = &find.ID
	}

	auth := entities.NewAuth(payload, find)
	if err := auth.Validate(); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRDOMAIN,
			Errors: multilerr,
		}
	}

	if find == nil {
		if err := x.authRepo.Create(ctx, auth); err != nil {
			multilerr = multierror.Append(multilerr, err)
			return nil, &exceptions.CustomError{
				Status: exceptions.ERRREPOSITORY,
				Errors: multilerr,
			}
		}

		// saves the default role for new users
		role, _ := x.roleRepo.FindByName(ctx, middleware.ROLE_MEMBER)
		if role != nil {
			payloadRoleUser := make([]*entities.RoleUser, 0)

			payloadRoleUser = append(payloadRoleUser, &entities.RoleUser{
				RoleId: role.ID.String(),
				UserId: auth.UserId,
			})

			if err := x.roleUserRepo.CreateMany(ctx, payloadRoleUser); err != nil {
				log.Println(err)
			}
		}
	} else {
		if err := x.authRepo.Update(ctx, auth); err != nil {
			multilerr = multierror.Append(multilerr, err)
			return nil, &exceptions.CustomError{
				Status: exceptions.ERRREPOSITORY,
				Errors: multilerr,
			}
		}
	}

	return auth, nil
}

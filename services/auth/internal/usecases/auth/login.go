package auth

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"

	"github.com/hashicorp/go-multierror"
)

func (x *authInteractor) Login(ctx context.Context, payload entities.AuthDto) (*entities.AuthTokenMeta, *exceptions.CustomError) {
	var multilerr *multierror.Error

	log.Println("start check email already")
	auth, err := x.authRepo.FindByEmail(ctx, payload.Email)
	if err != nil {
		multilerr = multierror.Append(multilerr, lang.ErrEmailNotFound)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	log.Println("start check same password hash")
	if !utils.CheckPasswordHash(payload.Password, auth.Password) {
		multilerr = multierror.Append(multilerr, lang.ErrPasswordIsIncorrent)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRDOMAIN,
			Errors: multilerr,
		}
	}

	log.Println("success login")

	return entities.NewAuthLogin(auth), nil
}

package auth

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"

	"github.com/hashicorp/go-multierror"
)

func (x *authInteractor) PasswordReset(ctx context.Context, payload entities.PasswordReset) *exceptions.CustomError {
	var multilerr *multierror.Error

	log.Println("PasswordReset::info#1:", "check if the token is valid")
	plaintext, err := utils.ChiperDecrypt(payload.Token, x.cfg.AppSecretKey)
	if err != nil {
		log.Println("PasswordReset::error#1:", err)
		multilerr = multierror.Append(multilerr, lang.TokenNotValid)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	data := strings.Split(fmt.Sprintf("%s", plaintext), ":")

	log.Println("PasswordReset::info#2:", "check if the token is expired")
	if utils.TimestampToTime(data[1]).Before(utils.TimeUTC()) {
		log.Println("PasswordReset::error#2:", err)
		multilerr = multierror.Append(multilerr, lang.TokenHasExpired)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	log.Println("PasswordReset::info#3:", "start check auth already")
	auth, err := x.authRepo.FindByUserId(ctx, data[0])
	if err != nil {
		log.Println("PasswordReset::error#3:", err)
		multilerr = multierror.Append(multilerr, lang.ErrEmailNotFound)
		return &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	// set new password
	auth.SetPasswordHash(payload.Password)

	log.Println("PasswordReset::info#4:", "update auth data")
	if err := x.authRepo.Update(ctx, auth); err != nil {
		log.Println("PasswordReset::error#4:", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	log.Println("PasswordReset::success#1:", "password reset")

	return nil
}

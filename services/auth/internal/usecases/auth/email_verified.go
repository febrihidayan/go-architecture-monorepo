package auth

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"

	"github.com/hashicorp/go-multierror"
)

func (x *authInteractor) EmailVerified(ctx context.Context, token string) *exceptions.CustomError {
	var multilerr *multierror.Error

	log.Println("EmailVerified::info#1:", "check if the token is valid")
	plaintext, err := utils.ChiperDecrypt(token, x.cfg.AppSecretKey)
	if err != nil {
		log.Println("EmailVerified::error#1:", err)
		multilerr = multierror.Append(multilerr, lang.TokenNotValid)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	payload := strings.Split(fmt.Sprintf("%s", plaintext), ":")

	log.Println("EmailVerified::info#2:", "check if the token is expired")
	if utils.TimestampToTime(payload[1]).Before(utils.TimeUTC()) {
		log.Println("EmailVerified::error#2:", err)
		multilerr = multierror.Append(multilerr, lang.TokenHasExpired)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	log.Println("EmailVerified::info#3:", "start check email already")
	auth, err := x.authRepo.FindByEmail(ctx, payload[0])
	if err != nil {
		log.Println("EmailVerified::error#3:", err)
		multilerr = multierror.Append(multilerr, lang.ErrEmailNotFound)
		return &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	log.Println("EmailVerified::info#4", "check if the email has been verified")
	if !auth.EmailVerifiedAt.IsZero() {
		log.Println("EmailVerified::error#4", err)
		multilerr = multierror.Append(multilerr, lang.EmailAddressVerified)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	// set email verified with time now
	auth.SetEmailVerifiedAt()

	log.Println("EmailVerified::info#5:", "update user data")
	if err := x.authRepo.Update(ctx, auth); err != nil {
		log.Println("EmailVerified::error#5:", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	// update email verified in user service
	go func(auth *entities.Auth) {
		id, _ := common.StringToID(auth.UserId)
		payload := entities.User{
			ID:              id,
			EmailVerifiedAt: auth.EmailVerifiedAt,
		}
		ctx := context.Background()
		if err := x.userRepo.UpdateEmailVerifiedUser(ctx, payload); err != nil {
			log.Println("EmailVerified::error#6:", err)

		}
	}(auth)

	log.Println("EmailVerified::success#1:", "email verified")

	return nil
}

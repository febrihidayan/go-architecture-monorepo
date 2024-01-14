package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"

	"github.com/hashicorp/go-multierror"
)

func (x *authInteractor) PasswordEmail(ctx context.Context, email string) *exceptions.CustomError {
	var (
		multilerr *multierror.Error
	)

	log.Println("PasswordEmail::info#1", "start check email already")
	auth, err := x.authRepo.FindByEmail(ctx, email)
	if err != nil {
		log.Println("PasswordEmail::error#1", err)
		multilerr = multierror.Append(multilerr, lang.ErrEmailNotFound)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	log.Println("PasswordEmail::info#2", "create encryption code")
	plainText := fmt.Sprintf("%s:%d", auth.UserId, utils.TimeUTC().Add(time.Hour*1).Unix())
	ciphertext, err := utils.ChiperEncrypt(plainText, x.cfg.AppSecretKey)
	if err != nil {
		log.Println("PasswordEmail::error#2", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	// send email password reset
	go func(auth *entities.Auth, token string) {
		ctx := context.Background()

		data := map[string]string{
			"link":   fmt.Sprintf("%s/auth/password/reset/%s", x.cfg.AppURL, token),
			"expire": "60", // 60 minutes
		}

		dataJson, err := json.Marshal(data)
		if err != nil {
			log.Println("PasswordEmail::error#3:", err)
		}

		payload := entities.NotificationSends{
			UserId:       auth.UserId,
			TemplateName: entities.TemplateTypePasswordReset,
			Data:         string(dataJson),
			Services:     []string{entities.NotificationTypeEmail},
			PathEmail:    "password-reset.html",
		}

		if err := x.notificationGrpcRepo.SendNotification(ctx, payload); err != nil {
			log.Println("PasswordEmail::error#4:", err)
		}
	}(auth, ciphertext)

	log.Println("PasswordEmail::success#1", "send reset link email")

	return nil
}

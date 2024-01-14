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

func (x *authInteractor) SendEmailVerified(ctx context.Context, email string) *exceptions.CustomError {
	var (
		multilerr *multierror.Error
	)

	log.Println("SendEmailVerified::info#1", "start check email already")
	auth, err := x.authRepo.FindByEmail(ctx, email)
	if err != nil {
		log.Println("SendEmailVerified::error#1", err)
		multilerr = multierror.Append(multilerr, lang.ErrEmailNotFound)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	log.Println("SendEmailVerified::info#2", "check if the email has been verified")
	if !auth.EmailVerifiedAt.IsZero() {
		log.Println("SendEmailVerified::error#2", err)
		multilerr = multierror.Append(multilerr, lang.EmailAddressVerified)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	log.Println("SendEmailVerified::info#3", "create encryption code")
	plainText := fmt.Sprintf("%s:%d", auth.Email, utils.TimeUTC().Add(time.Hour*1).Unix())
	ciphertext, err := utils.ChiperEncrypt(plainText, x.cfg.AppSecretKey)
	if err != nil {
		log.Println("SendEmailVerified::error#3", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	// send email register (welcome)
	go func(auth *entities.Auth, token string) {
		ctx := context.Background()

		data := map[string]string{
			"link": fmt.Sprintf("%s/auth/email/%s", x.cfg.AppURL, token),
		}

		dataJson, err := json.Marshal(data)
		if err != nil {
			log.Println("SendEmailVerified::error#4:", err)
		}

		payload := entities.NotificationSends{
			UserId:       auth.UserId,
			TemplateName: entities.TemplateTypeEmailVerified,
			Data:         string(dataJson),
			Services:     []string{entities.NotificationTypeEmail},
			PathEmail:    "email-verified.html",
		}

		if err := x.notificationGrpcRepo.SendNotification(ctx, payload); err != nil {
			log.Println("SendEmailVerified::error#5:", err)
		}
	}(auth, ciphertext)

	log.Println("SendEmailVerified::success#1", "send email verified")

	return nil
}

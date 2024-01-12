package auth

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/middleware"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"

	"github.com/hashicorp/go-multierror"
)

func (x *authInteractor) Register(ctx context.Context, payload entities.RegisterDto) (*entities.Auth, *exceptions.CustomError) {
	var multilerr *multierror.Error

	log.Println("Register::info#1:", "start check email already")

	find, _ := x.authRepo.FindByEmail(ctx, payload.Email)
	if find != nil {
		multilerr = multierror.Append(multilerr, errors.New("The email is already registered."))
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	log.Println("Register::info#2:", "start create user")

	user, err := x.userRepo.CreateUser(ctx, entities.User{
		Name:  payload.Name,
		Email: payload.Email,
	})
	if err != nil {
		log.Println("Register::error#1:", err)
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	log.Println("Register::info#3:", "start auth dto")

	auth := entities.NewAuth(entities.AuthDto{
		UserId:   user.ID.String(),
		Email:    payload.Email,
		Password: payload.Password,
	})

	log.Println("Register::info#4:", "start validation auth")

	if err := auth.Validate(); err != nil {
		log.Println("Register::error#2:", err)
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRDOMAIN,
			Errors: multilerr,
		}
	}

	log.Println("Register::info#5:", "start create auth")

	if err := x.authRepo.Create(ctx, auth); err != nil {
		log.Println("Register::error#3:", err)
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
			log.Println("Register::error#4:", err)
		}
	}

	// send email register (welcome)
	go func(auth *entities.Auth, user *entities.User) {
		ctx := context.Background()

		data := map[string]string{
			"name": user.Name,
		}

		dataJson, err := json.Marshal(data)
		if err != nil {
			log.Println("Register::error#5:", err)
		}

		payload := entities.NotificationSends{
			UserId:       auth.UserId,
			TemplateName: entities.TemplateTypeWelcome,
			Data:         string(dataJson),
			Services:     []string{entities.NotificationTypeEmail},
			PathEmail:    "welcome.html",
		}

		if err := x.notificationGrpcRepo.SendNotification(ctx, payload); err != nil {
			log.Println("Register::error#6:", err)
		}
	}(auth, user)

	log.Println("Register::success#1:", "success create user and auth")

	return auth, nil
}

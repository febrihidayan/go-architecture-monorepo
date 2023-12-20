package auth

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/middleware"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
)

func (x *AuthUsecaseSuite) TestLogin() {
	id := common.NewID()
	var auth *entities.Auth

	payloadDto := entities.AuthDto{
		Email:    "admin@app.com",
		Password: "password",
	}

	auth = &entities.Auth{
		ID:        id,
		Email:     "admin@app.com",
		Role:      middleware.ROLE_MEMBER,
		Password:  "$2a$10$bft5emRZGI1MWmAObIxmIO9Bs61t5P23pwjHq.5IBbEXTgLpIhdv.",
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	authMeta := entities.NewAuthLogin(auth)

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.authRepo.Mock.On("FindByEmail", auth.Email).Return(auth, nil)

				result, err := x.authUsecase.Login(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(result, authMeta)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.authRepo.Mock.On("FindByEmail", auth.Email).Return(nil, errors.New(mock.Anything))

				_, err := x.authUsecase.Login(context.Background(), payloadDto)

				e := &exceptions.CustomError{
					Status: exceptions.ERRREPOSITORY,
					Errors: multierror.Append(lang.ErrEmailNotFound),
				}

				x.Equal(err, e)
			},
		},
		{
			name: "Password Negatif Case",
			tests: func(arg Any) {
				x.authRepo.Mock.On("FindByEmail", auth.Email).Return(auth, nil)

				payloadDto.Password = "wordpass"
				_, err := x.authUsecase.Login(context.Background(), payloadDto)

				e := &exceptions.CustomError{
					Status: exceptions.ERRDOMAIN,
					Errors: multierror.Append(lang.ErrPasswordIsIncorrent),
				}

				x.Equal(err, e)
			},
		},
	}

	for _, arg := range args {
		x.Run(arg.name, func() {
			x.SetupTest()
			arg.tests(arg.args)
		})
	}
}

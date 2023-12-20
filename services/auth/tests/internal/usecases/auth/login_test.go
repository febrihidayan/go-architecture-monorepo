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
	var (
		id        = common.NewID()
		auth      *entities.Auth
		roleUsers []*entities.RoleUser
		role      *entities.Role
	)

	payloadDto := entities.AuthDto{
		Email:    "admin@app.com",
		Password: "password",
	}

	auth = &entities.Auth{
		ID:        id,
		Email:     "admin@app.com",
		Password:  "$2a$10$bft5emRZGI1MWmAObIxmIO9Bs61t5P23pwjHq.5IBbEXTgLpIhdv.",
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	authMeta := entities.NewAuthLogin(&entities.AuthMeta{
		Auth:  auth,
		Roles: []string{middleware.ROLE_SUPERADMINISTRATOR},
	})

	role = &entities.Role{
		ID:          id,
		Name:        middleware.ROLE_SUPERADMINISTRATOR,
		DisplayName: middleware.ROLE_SUPERADMINISTRATOR,
		Description: middleware.ROLE_SUPERADMINISTRATOR,
		CreatedAt:   utils.TimeUTC(),
		UpdatedAt:   utils.TimeUTC(),
	}

	roleUsers = append(roleUsers, &entities.RoleUser{
		UserId: id.String(),
		RoleId: id.String(),
	})

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.authRepo.Mock.On("FindByEmail", auth.Email).Return(auth, nil)

				x.roleUserRepo.Mock.On("AllByUserId", auth.UserId).Return(roleUsers, nil)

				for _, item := range roleUsers {
					x.roleRepo.Mock.On("Find", item.RoleId).Return(role, nil)
				}

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

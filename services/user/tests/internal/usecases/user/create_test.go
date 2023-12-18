package user

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
)

func (x *UserUsecaseSuite) TestCreate() {
	id := common.NewID()
	var user *entities.User

	payloadDto := entities.UserDto{
		ID:    &id,
		Name:  "Admin",
		Email: "admin@app.com",
		Role:  "admin",
	}

	user = &entities.User{
		ID:        id,
		Name:      "Admin",
		Email:     "admin@app.com",
		Role:      "admin",
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.userRepo.Mock.On("Create", user).Return(nil)

				result, err := x.userUsecase.Create(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(result, user)
			},
		},
		{
			name: "Failed Negative Case",
			tests: func(arg Any) {
				x.userRepo.Mock.On("Create", user).Return(errors.New(mock.Anything))

				_, err := x.userUsecase.Create(context.Background(), payloadDto)

				e := &exceptions.CustomError{
					Status: exceptions.ERRREPOSITORY,
					Errors: multierror.Append(errors.New(mock.Anything)),
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

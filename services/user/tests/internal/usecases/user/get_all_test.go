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

func (x *UserUsecaseSuite) TestGetAll() {
	id := common.NewID()
	var users []*entities.User

	users = append(users, &entities.User{
		ID:        id,
		Name:      "Admin",
		Email:     "admin@app.com",
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	})

	params := entities.UserQueryParams{
		Search:  "",
		Page:    1,
		PerPage: 10,
	}

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.userRepo.Mock.On("GetAll", &params).Return(users, len(users), nil)

				results, err := x.userUsecase.GetAll(context.Background(), params)
				x.Nil(err)
				x.Equal(results, &entities.UserMeta{
					Data:  users,
					Total: len(users),
				})
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.userRepo.Mock.On("GetAll", &params).Return(nil, nil, errors.New(mock.Anything))

				_, err := x.userUsecase.GetAll(context.Background(), params)

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

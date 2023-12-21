package profile

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

func (x *ProfileUsecaseSuite) TestFind() {
	id := common.NewID()
	var user *entities.User

	user = &entities.User{
		ID:        id,
		Name:      "Admin",
		Email:     "admin@app.com",
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
				x.userRepo.Mock.On("Find", id.String()).Return(user, nil)

				profile, err := x.profileUsecase.Find(context.Background(), id.String())
				x.Nil(err)
				x.Equal(profile, user)
			},
		},
		{
			name: "Failed Negative Case",
			tests: func(arg Any) {
				x.userRepo.Mock.On("Find", id.String()).Return(nil, errors.New(mock.Anything))

				_, err := x.profileUsecase.Find(context.Background(), id.String())

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

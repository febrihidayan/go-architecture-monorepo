package profile

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/stretchr/testify/mock"
)

func (x *ProfileUsecaseSuite) TestUpdate() {
	id := common.NewID()
	var user *entities.User

	payloadDto := entities.UserDto{
		ID:    &id,
		Name:  "Admin",
		Email: "admin@app.com",
	}

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
				x.userRepo.Mock.On("Find", payloadDto.ID.String()).Return(user, nil)

				x.userRepo.Mock.On("Update", user).Return(nil)

				result, err := x.profileUsecase.Update(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(user, result)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.userRepo.Mock.On("Find", payloadDto.ID.String()).Return(nil, errors.New(mock.Anything))

				_, err := x.profileUsecase.Update(context.Background(), payloadDto)
				x.NotNil(err)
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

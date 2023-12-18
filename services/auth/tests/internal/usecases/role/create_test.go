package role

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

func (x *RoleUsecaseSuite) TestCreate() {
	id := common.NewID()
	var role *entities.Role

	payloadDto := entities.RoleDto{
		ID:          &id,
		Name:        "superadmin",
		DisplayName: "Super Admin",
		Description: "Super Admin",
	}

	role = &entities.Role{
		ID:          id,
		Name:        "superadmin",
		DisplayName: "Super Admin",
		Description: "Super Admin",
		CreatedAt:   utils.TimeUTC(),
		UpdatedAt:   utils.TimeUTC(),
	}

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.roleRepo.Mock.On("FindByName", payloadDto.Name).Return(nil, errors.New(mock.Anything))

				x.roleRepo.Mock.On("Create", role).Return(nil)

				result, err := x.roleUsecase.Create(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(role, result)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.roleRepo.Mock.On("FindByName", payloadDto.Name).Return(nil, errors.New(mock.Anything))

				x.roleRepo.Mock.On("Create", role).Return(errors.New(mock.Anything))

				_, err := x.roleUsecase.Create(context.Background(), payloadDto)
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

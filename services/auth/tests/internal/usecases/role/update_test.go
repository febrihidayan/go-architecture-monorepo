package role

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

func (x *RoleUsecaseSuite) TestUpdate() {
	id := common.NewID()
	var role *entities.Role
	var roleChangeName *entities.Role

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

	roleChangeName = &entities.Role{
		ID:          id,
		Name:        "admin",
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
				x.roleRepo.Mock.On("Find", payloadDto.ID.String()).Return(role, nil)

				x.roleRepo.Mock.On("Update", role).Return(nil)

				result, err := x.roleUsecase.Update(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(role, result)
			},
		},
		{
			name: "Success Change Name Positive Case",
			tests: func(arg Any) {
				x.roleRepo.Mock.On("Find", payloadDto.ID.String()).Return(role, nil)

				// change name
				payloadDto.Name = "admin"

				x.roleRepo.Mock.On("FindByName", payloadDto.Name).Return(nil, errors.New(mock.Anything))

				x.roleRepo.Mock.On("Update", roleChangeName).Return(nil)

				result, err := x.roleUsecase.Update(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(roleChangeName, result)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.roleRepo.Mock.On("Find", payloadDto.ID.String()).Return(nil, errors.New(mock.Anything))

				_, err := x.roleUsecase.Update(context.Background(), payloadDto)
				x.NotNil(err)
			},
		},
		{
			name: "Failed Change Name Negatif Case",
			tests: func(arg Any) {
				x.roleRepo.Mock.On("Find", payloadDto.ID.String()).Return(role, nil)

				// change name
				payloadDto.Name = "admin"

				x.roleRepo.Mock.On("FindByName", payloadDto.Name).Return(roleChangeName, nil)

				_, err := x.roleUsecase.Update(context.Background(), payloadDto)
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

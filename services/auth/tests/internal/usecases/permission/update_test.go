package permission

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

func (x *PermissionUsecaseSuite) TestUpdate() {
	id := common.NewID()
	var permission *entities.Permission
	var permissionChangeName *entities.Permission

	payloadDto := entities.PermissionDto{
		ID:          &id,
		Name:        "users_create",
		DisplayName: "Users Create",
		Description: "Users Create",
	}

	permission = &entities.Permission{
		ID:          id,
		Name:        "users_create",
		DisplayName: "Users Create",
		Description: "Users Create",
		CreatedAt:   utils.TimeUTC(),
		UpdatedAt:   utils.TimeUTC(),
	}

	permissionChangeName = &entities.Permission{
		ID:          id,
		Name:        "users_update",
		DisplayName: "Users Create",
		Description: "Users Create",
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
				x.permissionRepo.Mock.On("Find", payloadDto.ID.String()).Return(permission, nil)

				x.permissionRepo.Mock.On("Update", permission).Return(nil)

				result, err := x.permissionUsecase.Update(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(permission, result)
			},
		},
		{
			name: "Success Change Name Positive Case",
			tests: func(arg Any) {
				x.permissionRepo.Mock.On("Find", payloadDto.ID.String()).Return(permission, nil)

				// change name
				payloadDto.Name = "users_update"

				x.permissionRepo.Mock.On("FindByName", payloadDto.Name).Return(nil, errors.New(mock.Anything))

				x.permissionRepo.Mock.On("Update", permissionChangeName).Return(nil)

				result, err := x.permissionUsecase.Update(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(permissionChangeName, result)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.permissionRepo.Mock.On("Find", payloadDto.ID.String()).Return(nil, errors.New(mock.Anything))

				_, err := x.permissionUsecase.Update(context.Background(), payloadDto)
				x.NotNil(err)
			},
		},
		{
			name: "Failed Change Name Negatif Case",
			tests: func(arg Any) {
				x.permissionRepo.Mock.On("Find", payloadDto.ID.String()).Return(permission, nil)

				// change name
				payloadDto.Name = "users_update"

				x.permissionRepo.Mock.On("FindByName", payloadDto.Name).Return(permissionChangeName, nil)

				_, err := x.permissionUsecase.Update(context.Background(), payloadDto)
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

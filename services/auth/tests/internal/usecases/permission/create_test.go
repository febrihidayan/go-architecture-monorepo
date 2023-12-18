package permission

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/stretchr/testify/mock"
)

func (x *PermissionUsecaseSuite) TestCreate() {
	id := common.NewID()
	var permission *entities.Permission

	payloadDto := entities.PermissionDto{
		ID:          &id,
		Name:        "superadmin",
		DisplayName: "Super Admin",
		Description: "Super Admin",
	}

	permission = &entities.Permission{
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
				x.permissionRepo.Mock.On("FindByName", payloadDto.Name).Return(nil, errors.New(mock.Anything))

				x.permissionRepo.Mock.On("Create", permission).Return(nil)

				result, err := x.permissionUsecase.Create(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(permission, result)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.permissionRepo.Mock.On("FindByName", payloadDto.Name).Return(nil, errors.New(mock.Anything))

				x.permissionRepo.Mock.On("Create", permission).Return(errors.New(mock.Anything))

				_, err := x.permissionUsecase.Create(context.Background(), payloadDto)
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

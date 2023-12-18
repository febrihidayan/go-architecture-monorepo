package permission

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
)

func (x *PermissionUsecaseSuite) TestFind() {
	id := common.NewID()
	var permission *entities.Permission

	permission = &entities.Permission{
		ID:          id,
		Name:        "users_create",
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
				x.permissionRepo.Mock.On("Find", id.String()).Return(permission, nil)

				result, err := x.permissionUsecase.Find(context.Background(), id.String())
				x.Nil(err)
				x.Equal(result, permission)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.permissionRepo.Mock.On("Find", id.String()).Return(nil, errors.New(mock.Anything))

				_, err := x.permissionUsecase.Find(context.Background(), id.String())

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

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

func (x *PermissionUsecaseSuite) TestGetAll() {
	id := common.NewID()
	var permissions []*entities.Permission

	permissions = append(permissions, &entities.Permission{
		ID:          id,
		Name:        "users_create",
		DisplayName: "Users Create",
		Description: "Users Create",
		CreatedAt:   utils.TimeUTC(),
		UpdatedAt:   utils.TimeUTC(),
	})

	params := entities.PermissionQueryParams{
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
				x.permissionRepo.Mock.On("GetAll", &params).Return(permissions, len(permissions), nil)

				results, err := x.permissionUsecase.GetAll(context.Background(), params)
				x.Nil(err)
				x.Equal(results, &entities.PermissionMeta{
					Data:  permissions,
					Total: len(permissions),
				})
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.permissionRepo.Mock.On("GetAll", &params).Return(nil, nil, errors.New(mock.Anything))

				_, err := x.permissionUsecase.GetAll(context.Background(), params)

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

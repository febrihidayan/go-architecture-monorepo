package role

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

func (x *RoleUsecaseSuite) TestGetAll() {
	id := common.NewID()
	var roles []*entities.Role

	roles = append(roles, &entities.Role{
		ID:          id,
		Name:        "superadmin",
		DisplayName: "Super Admin",
		Description: "Super Admin",
		CreatedAt:   utils.TimeUTC(),
		UpdatedAt:   utils.TimeUTC(),
	})

	params := entities.RoleQueryParams{
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
				x.roleRepo.Mock.On("GetAll", &params).Return(roles, len(roles), nil)

				results, err := x.roleUsecase.GetAll(context.Background(), params)
				x.Nil(err)
				x.Equal(results, &entities.RoleMeta{
					Data:  roles,
					Total: len(roles),
				})
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.roleRepo.Mock.On("GetAll", &params).Return(nil, nil, errors.New(mock.Anything))

				_, err := x.roleUsecase.GetAll(context.Background(), params)

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

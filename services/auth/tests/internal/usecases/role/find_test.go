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

func (x *RoleUsecaseSuite) TestFind() {
	id := common.NewID()
	var role *entities.Role

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
				x.roleRepo.Mock.On("Find", id.String()).Return(role, nil)

				result, err := x.roleUsecase.Find(context.Background(), id.String())
				x.Nil(err)
				x.Equal(result, role)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.roleRepo.Mock.On("Find", id.String()).Return(nil, errors.New(mock.Anything))

				_, err := x.roleUsecase.Find(context.Background(), id.String())

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

package acl

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
)

func (x *AclUsecaseSuite) TestGetAllRole() {
	var (
		roles []*entities.Role
	)

	roles = append(roles, x.role)

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.roleRepo.Mock.On("All").Return(roles, nil)

				result, err := x.aclUsecase.GetAllRole(context.Background())
				x.Nil(err)
				x.Equal(result, roles)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.roleRepo.Mock.On("All").Return(nil, errors.New(mock.Anything))

				_, err := x.aclUsecase.GetAllRole(context.Background())

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

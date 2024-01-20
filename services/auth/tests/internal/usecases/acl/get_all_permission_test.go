package acl

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
)

func (x *AclUsecaseSuite) TestGetAllPermission() {
	var (
		permissions []*entities.Permission
	)

	permissions = append(permissions, x.permission)

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.permissionRepo.Mock.On("All").Return(permissions, nil)

				result, err := x.aclUsecase.GetAllPermission(context.Background())
				x.Nil(err)
				x.Equal(result, permissions)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.permissionRepo.Mock.On("All").Return(nil, errors.New(mock.Anything))

				_, err := x.aclUsecase.GetAllPermission(context.Background())

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

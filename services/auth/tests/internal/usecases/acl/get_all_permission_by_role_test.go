package acl

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
)

func (x *AclUsecaseSuite) TestGetAllPermissionByRole() {
	var (
		permissionRoles []*entities.PermissionRole
		permissions     []*entities.Permission
	)

	permissionRoles = append(permissionRoles, x.permissionRole)
	permissions = append(permissions, x.permission)

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.permissionRoleRepo.Mock.On("AllByRoleId", x.id.String()).Return(permissionRoles, nil)

				x.permissionRepo.Mock.On("Find", x.permissionRole.PermissionId).Return(x.permission, nil)

				result, err := x.aclUsecase.GetAllPermissionByRole(context.Background(), x.id.String())
				x.Nil(err)
				x.Equal(result, permissions)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.permissionRoleRepo.Mock.On("AllByRoleId", x.id.String()).Return(nil, errors.New(mock.Anything))

				_, err := x.aclUsecase.GetAllPermissionByRole(context.Background(), x.id.String())

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

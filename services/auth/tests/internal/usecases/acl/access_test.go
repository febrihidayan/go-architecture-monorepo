package acl

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
)

func (x *AclUsecaseSuite) TestAccess() {
	var (
		roles       []*entities.Role
		permissions []*entities.Permission
	)

	roles = append(roles, x.role)
	permissions = append(permissions, x.permission)

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.authRepo.Mock.On("FindByUserId", x.id.String()).Return(x.auth, nil)
				x.roleRepo.Mock.On("AllByUserId", x.id.String()).Return(roles, nil)
				x.permissionRepo.Mock.On("AllPermissionByUserId", x.id.String()).Return(permissions, nil)

				result, err := x.aclUsecase.AccessUserLogin(context.Background(), x.id.String())
				x.Nil(err)
				x.Equal(result, &entities.AclMeta{
					Permissions: permissions,
					Roles:       roles,
				})
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.authRepo.Mock.On("FindByUserId", x.id.String()).Return(nil, errors.New(mock.Anything))

				_, err := x.aclUsecase.AccessUserLogin(context.Background(), x.id.String())

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

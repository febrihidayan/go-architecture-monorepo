package acl

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
)

func (x *AclUsecaseSuite) TestUpdatePermissionByRole() {
	var (
		newID              = common.NewID()
		newPermissionRoles []*entities.PermissionRole
		permissionRoles    []*entities.PermissionRole
		deletes            []string
	)

	payload := entities.AclPermissionDto{
		RoleId: x.role.ID.String(),
		Permissions: []string{
			x.id.String(),
			newID.String(),
		},
	}

	payloadDeleteAll := entities.AclPermissionDto{
		RoleId:      x.role.ID.String(),
		Permissions: []string{},
	}

	deletes = []string{
		x.id.String(),
	}

	newPermissionRoles = append(newPermissionRoles, entities.NewPermissionRole(entities.PermissionRoleDto{
		PermissionId: newID.String(),
		RoleId:       x.id.String(),
	}))

	permissionRoles = append(permissionRoles, x.permissionRole)

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.roleRepo.Mock.On("Find", x.id.String()).Return(x.role, nil)
				x.permissionRoleRepo.Mock.On("AllByRoleId", x.role.ID.String()).Return(permissionRoles, nil)
				x.permissionRoleRepo.Mock.On("CreateMany", newPermissionRoles).Return(nil)
				x.permissionRoleRepo.Mock.On("DeleteByPermissionIds", deletes).Return(nil)

				err := x.aclUsecase.UpdatePermissionByRole(context.Background(), payload)
				x.Nil(err)
			},
		},
		{
			name: "Success Delete Permissions Positive Case",
			tests: func(arg Any) {
				x.roleRepo.Mock.On("Find", x.id.String()).Return(x.role, nil)
				x.permissionRoleRepo.Mock.On("AllByRoleId", x.role.ID.String()).Return(permissionRoles, nil)
				x.permissionRoleRepo.Mock.On("DeleteByRoleId", x.role.ID.String()).Return(nil)

				err := x.aclUsecase.UpdatePermissionByRole(context.Background(), payloadDeleteAll)
				x.Nil(err)
			},
		},
		{
			name: "Failed Find Role Negatif Case",
			tests: func(arg Any) {
				x.roleRepo.Mock.On("Find", x.id.String()).Return(nil, errors.New(mock.Anything))

				err := x.aclUsecase.UpdatePermissionByRole(context.Background(), payload)

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

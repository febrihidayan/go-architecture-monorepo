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

func (x *AclUsecaseSuite) TestUpdateUser() {
	var (
		newID              = common.NewID()
		newPermissionUsers []*entities.PermissionUser
		newRoleUsers       []*entities.RoleUser
		permissionUsers    []*entities.PermissionUser
		roleUsers          []*entities.RoleUser
		emptyDelete        []string
	)

	payload := entities.AclUserDto{
		UserId: x.id.String(),
		Permissions: []string{
			x.id.String(),
			newID.String(),
		},
		Roles: []string{
			x.id.String(),
			newID.String(),
		},
	}

	payloadChange := entities.AclUserDto{
		UserId: x.id.String(),
		Permissions: []string{
			newID.String(),
		},
		Roles: []string{
			newID.String(),
		},
	}

	payloadDeleteAll := entities.AclUserDto{
		UserId:      x.id.String(),
		Permissions: []string{},
		Roles:       []string{},
	}

	deleteChange := []string{
		x.id.String(),
	}

	newPermissionUsers = append(newPermissionUsers, entities.NewPermissionUser(entities.PermissionUserDto{
		PermissionId: newID.String(),
		UserId:       x.id.String(),
	}))

	newRoleUsers = append(newRoleUsers, entities.NewRoleUser(entities.RoleUserDto{
		RoleId: newID.String(),
		UserId: x.id.String(),
	}))

	permissionUsers = append(permissionUsers, x.permissionUser)
	roleUsers = append(roleUsers, x.roleUser)

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.authRepo.Mock.On("FindByUserId", x.id.String()).Return(x.auth, nil)
				x.roleUserRepo.Mock.On("AllByUserId", x.auth.UserId).Return(roleUsers, nil)
				x.permissionUserRepo.Mock.On("AllByUserId", x.auth.UserId).Return(permissionUsers, nil)

				// permission
				x.permissionUserRepo.Mock.On("CreateMany", newPermissionUsers).Return(nil)
				x.permissionUserRepo.Mock.On("DeleteByPermissionIds", emptyDelete).Return(nil)
				x.permissionUserRepo.Mock.On("DeleteByUserId", x.auth.UserId).Return(nil)

				// role
				x.roleUserRepo.Mock.On("CreateMany", newRoleUsers).Return(nil)
				x.roleUserRepo.Mock.On("DeleteByRoleIds", emptyDelete).Return(nil)
				x.roleUserRepo.Mock.On("DeleteByUserId", x.auth.UserId).Return(nil)

				err := x.aclUsecase.UpdateUser(context.Background(), payload)
				x.Nil(err)
			},
		},
		{
			name: "Success Change Permissions and Roles Positive Case",
			tests: func(arg Any) {
				x.authRepo.Mock.On("FindByUserId", x.id.String()).Return(x.auth, nil)
				x.roleUserRepo.Mock.On("AllByUserId", x.auth.UserId).Return(roleUsers, nil)
				x.permissionUserRepo.Mock.On("AllByUserId", x.auth.UserId).Return(permissionUsers, nil)

				// permission
				x.permissionUserRepo.Mock.On("CreateMany", newPermissionUsers).Return(nil)
				x.permissionUserRepo.Mock.On("DeleteByPermissionIds", deleteChange).Return(nil)
				x.permissionUserRepo.Mock.On("DeleteByUserId", x.auth.UserId).Return(nil)

				// role
				x.roleUserRepo.Mock.On("CreateMany", newRoleUsers).Return(nil)
				x.roleUserRepo.Mock.On("DeleteByRoleIds", deleteChange).Return(nil)
				x.roleUserRepo.Mock.On("DeleteByUserId", x.auth.UserId).Return(nil)

				err := x.aclUsecase.UpdateUser(context.Background(), payloadChange)
				x.Nil(err)
			},
		},
		{
			name: "Success Delete Permissions and Roles Positive Case",
			tests: func(arg Any) {
				x.authRepo.Mock.On("FindByUserId", x.id.String()).Return(x.auth, nil)
				x.roleUserRepo.Mock.On("AllByUserId", x.auth.UserId).Return(roleUsers, nil)
				x.permissionUserRepo.Mock.On("AllByUserId", x.auth.UserId).Return(permissionUsers, nil)

				// permission and delete
				x.permissionUserRepo.Mock.On("DeleteByUserId", x.auth.UserId).Return(nil)
				x.roleUserRepo.Mock.On("DeleteByUserId", x.auth.UserId).Return(nil)

				err := x.aclUsecase.UpdateUser(context.Background(), payloadDeleteAll)
				x.Nil(err)
			},
		},
		{
			name: "Failed Find Auth Negatif Case",
			tests: func(arg Any) {
				x.authRepo.Mock.On("FindByUserId", x.id.String()).Return(nil, errors.New(mock.Anything))

				err := x.aclUsecase.UpdateUser(context.Background(), payload)

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

package entities

import (
	"testing"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

func TestValidatePermissionUser(t *testing.T) {
	var uuid = common.NewID().String()
	var errs *multierror.Error

	tests := []struct {
		name string
		args entities.PermissionUserDto
		errs *multierror.Error
	}{
		{
			name: "Success Positive Case",
			args: entities.PermissionUserDto{
				PermissionId: uuid,
				UserId:       uuid,
			},
			errs: multierror.Append(errs, errs.ErrorOrNil()),
		},
		{
			name: "Failed Negative Case",
			args: entities.PermissionUserDto{
				PermissionId: "",
				UserId:       "",
			},
			errs: multierror.Append(errs, lang.Trans("filled", "PermissionId"), lang.Trans("filled", "UserId")),
		},
	}

	for _, test := range tests {
		args := entities.NewPermissionUser(test.args)

		if err := args.Validate(); err != nil {
			assert.ElementsMatch(t, err.Errors, test.errs.Errors, test.name)
		} else {
			assert.Nil(t, err)
		}
	}
}

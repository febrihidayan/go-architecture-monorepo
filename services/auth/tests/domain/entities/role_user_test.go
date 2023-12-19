package entities

import (
	"testing"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

func TestValidateRoleUser(t *testing.T) {
	var uuid = common.NewID().String()
	var errs *multierror.Error

	tests := []struct {
		name string
		args entities.RoleUserDto
		errs *multierror.Error
	}{
		{
			name: "Success Positive Case",
			args: entities.RoleUserDto{
				UserId: uuid,
				RoleId: uuid,
			},
			errs: multierror.Append(errs, errs.ErrorOrNil()),
		},
		{
			name: "Failed Negative Case",
			args: entities.RoleUserDto{
				UserId: "",
				RoleId: "",
			},
			errs: multierror.Append(errs, lang.Trans("filled", "UserId"), lang.Trans("filled", "RoleId")),
		},
	}

	for _, test := range tests {
		args := entities.NewRoleUser(test.args)

		if err := args.Validate(); err != nil {
			assert.ElementsMatch(t, err.Errors, test.errs.Errors, test.name)
		} else {
			assert.Nil(t, err)
		}
	}
}

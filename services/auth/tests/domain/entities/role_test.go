package entities

import (
	"testing"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

func TestValidateRole(t *testing.T) {
	var uuid = common.NewID()
	var errs *multierror.Error

	tests := []struct {
		name string
		args entities.RoleDto
		errs *multierror.Error
	}{
		{
			name: "Success Positive Case",
			args: entities.RoleDto{
				ID:          &uuid,
				Name:        "admin",
				DisplayName: "Admin",
			},
			errs: multierror.Append(errs, errs.ErrorOrNil()),
		},
		{
			name: "Failed Negative Case",
			args: entities.RoleDto{
				ID:          &uuid,
				Name:        "admin",
				DisplayName: "",
			},
			errs: multierror.Append(errs, lang.Trans("filled", "DisplayName")),
		},
	}

	for _, test := range tests {
		args := entities.NewRole(test.args)

		if err := args.Validate(); err != nil {
			assert.ElementsMatch(t, err.Errors, test.errs.Errors, test.name)
		} else {
			assert.Nil(t, err)
		}
	}
}

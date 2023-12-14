package entities

import (
	"testing"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"

	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
)

func TestValidateUser(t *testing.T) {
	var uuid = common.NewID()
	var errs *multierror.Error

	tests := []struct {
		name string
		args entities.UserDto
		errs *multierror.Error
	}{
		{
			name: "Success Positive Case",
			args: entities.UserDto{
				ID:    &uuid,
				Name:  "admin",
				Email: "admin@app.com",
				Role:  "admin",
			},
			errs: multierror.Append(errs, errs.ErrorOrNil()),
		},
		{
			name: "error not adding name and email",
			args: entities.UserDto{
				ID:    &uuid,
				Name:  "",
				Email: "",
				Role:  "admin",
			},
			errs: multierror.Append(errs, lang.ErrNameRequired, lang.ErrEmailRequired),
		},
		{
			name: "error not adding role",
			args: entities.UserDto{
				ID:    &uuid,
				Name:  "admin",
				Email: "admin@app.com",
				Role:  "",
			},
			errs: multierror.Append(errs, lang.ErrRoleRequired),
		},
	}

	for _, test := range tests {
		args := entities.NewUser(test.args)

		if err := args.Validate(); err != nil {
			assert.ElementsMatch(t, err.Errors, test.errs.Errors, test.name)
		} else {
			assert.Nil(t, err)
		}
	}
}

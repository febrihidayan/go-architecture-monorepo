package entities

import (
	"testing"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

func TestValidateAuth(t *testing.T) {
	var uuid = common.NewID()
	var UserId = common.NewID()
	var errs *multierror.Error

	tests := []struct {
		name string
		args entities.AuthDto
		errs *multierror.Error
	}{
		{
			name: "Success Positive Case",
			args: entities.AuthDto{
				ID:       &uuid,
				UserId:   UserId.String(),
				Email:    "admin@app.com",
				Password: "password",
				Role:     "admin",
			},
			errs: multierror.Append(errs, errs.ErrorOrNil()),
		},
		{
			name: "error not adding user_id",
			args: entities.AuthDto{
				ID:       &uuid,
				UserId:   "",
				Email:    "admin@app.com",
				Password: "password",
				Role:     "admin",
			},
			errs: multierror.Append(errs, lang.ErrUserIdRequired),
		},
		{
			name: "error not adding email and password",
			args: entities.AuthDto{
				ID:       &uuid,
				UserId:   UserId.String(),
				Email:    "",
				Password: "",
				Role:     "admin",
			},
			errs: multierror.Append(errs, lang.ErrEmailRequired, lang.ErrPasswordRequired),
		},
		{
			name: "error not adding role",
			args: entities.AuthDto{
				ID:       &uuid,
				UserId:   UserId.String(),
				Email:    "admin@app.com",
				Password: "password",
			},
			errs: multierror.Append(errs, lang.ErrRoleRequired),
		},
	}

	for _, test := range tests {
		args := entities.NewAuth(test.args)

		if err := args.Validate(); err != nil {
			assert.ElementsMatch(t, err.Errors, test.errs.Errors, test.name)
		} else {
			assert.Nil(t, err)
		}
	}
}

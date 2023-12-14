package entities

import (
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

func TestValidateRegister(t *testing.T) {
	var errs *multierror.Error

	tests := []struct {
		name string
		args entities.RegisterDto
		errs *multierror.Error
	}{
		{
			name: "Success Positive Case",
			args: entities.RegisterDto{
				Name:            "Admin",
				Email:           "admin@app.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			errs: multierror.Append(errs, errs.ErrorOrNil()),
		},
		{
			name: "error not adding name and email",
			args: entities.RegisterDto{
				Name:            "",
				Email:           "",
				Password:        "password",
				ConfirmPassword: "password",
			},
			errs: multierror.Append(errs, lang.ErrNameRequired, lang.ErrEmailRequired),
		},
		{
			name: "error not adding password and confirm password",
			args: entities.RegisterDto{
				Name:            "Admin",
				Email:           "admin@app.com",
				Password:        "",
				ConfirmPassword: "",
			},
			errs: multierror.Append(errs, lang.ErrPasswordRequired, lang.ErrConfirmPasswordRequired),
		},
		{
			name: "error confirmation passowrd is not the same",
			args: entities.RegisterDto{
				Name:            "Admin",
				Email:           "admin@app.com",
				Password:        "password",
				ConfirmPassword: "wordpass",
			},
			errs: multierror.Append(errs, lang.ErrConfitmPasswordNotSame),
		},
	}

	for _, test := range tests {
		args := entities.NewRegister(test.args)

		if err := args.Validate(); err != nil {
			assert.ElementsMatch(t, err.Errors, test.errs.Errors, test.name)
		} else {
			assert.Nil(t, err)
		}
	}
}

package entities

import (
	"testing"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

func TestValidateTemplate(t *testing.T) {
	var uuid = common.NewID()
	var errs *multierror.Error

	tests := []struct {
		name string
		args entities.TemplateDto
		errs *multierror.Error
	}{
		{
			name: "Success Positive Case",
			args: entities.TemplateDto{
				ID:   &uuid,
				Name: "register_user",
				Type: "in-app",
				Data: `{"title":{"en":"Welcome","id":"Selamat Datang"}}`,
			},
			errs: multierror.Append(errs, errs.ErrorOrNil()),
		},
		{
			name: "Failed Negative Case",
			args: entities.TemplateDto{
				ID:   &uuid,
				Name: "register_user",
				Type: "in-app",
				Data: "",
			},
			errs: multierror.Append(errs, lang.Trans("filled", "Data")),
		},
	}

	for _, test := range tests {
		args := entities.NewTemplate(test.args)

		if err := args.Validate(); err != nil {
			assert.ElementsMatch(t, err.Errors, test.errs.Errors, test.name)
		} else {
			assert.Nil(t, err)
		}
	}
}

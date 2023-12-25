package entities

import (
	"testing"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"

	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
)

func TestValidateCloud(t *testing.T) {
	var uuid = common.NewID()
	var errs *multierror.Error

	tests := []struct {
		name string
		args entities.CloudDto
		errs *multierror.Error
	}{
		{
			name: "Success Positive Case",
			args: entities.CloudDto{
				ID:        &uuid,
				Name:      "file.jpg",
				Url:       "http://domain.com/file.jpg",
				CreatedBy: uuid.String(),
			},
			errs: multierror.Append(errs, errs.ErrorOrNil()),
		},
		{
			name: "Failed Negative Case",
			args: entities.CloudDto{
				ID:        &uuid,
				Name:      "file.jpg",
				Url:       "",
				CreatedBy: uuid.String(),
			},
			errs: multierror.Append(errs, lang.Trans("filled", "Url")),
		},
	}

	for _, test := range tests {
		args := entities.NewCloud(test.args)

		if err := args.Validate(); err != nil {
			assert.ElementsMatch(t, err.Errors, test.errs.Errors, test.name)
		} else {
			assert.Nil(t, err)
		}
	}
}

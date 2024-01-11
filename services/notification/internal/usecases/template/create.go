package template

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *templateInteractor) Create(ctx context.Context, payload entities.TemplateDto) (*entities.Template, *exceptions.CustomError) {
	var multilerr *multierror.Error

	find, _ := x.templateRepo.FindByName(ctx, payload.Name)
	if find != nil {
		multilerr = multierror.Append(multilerr, lang.ErrTemplateAlready)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	template := entities.NewTemplate(payload)
	if err := template.Validate(); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRDOMAIN,
			Errors: multilerr,
		}
	}

	if err := x.templateRepo.Create(ctx, template); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return template, nil
}

package cloud

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *cloudInteractor) Deletes(ctx context.Context, payloads []*entities.Cloud) *exceptions.CustomError {
	var multilerr *multierror.Error

	for _, item := range payloads {
		find, err := x.cloudRepo.FindByUrl(ctx, item.Url)
		if err != nil {
			continue
		}

		if err := x.cloudRepo.Delete(ctx, find.ID.String()); err != nil {
			multilerr = multierror.Append(multilerr, err)
			return &exceptions.CustomError{
				Status: exceptions.ERRSYSTEM,
				Errors: multilerr,
			}
		}
	}

	return nil
}

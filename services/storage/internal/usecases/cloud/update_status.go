package cloud

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *cloudInteractor) UpdateStatus(ctx context.Context, payloads []*entities.Cloud) *exceptions.CustomError {
	var multilerr *multierror.Error

	for _, item := range payloads {
		find, _ := x.cloudRepo.Find(ctx, item.ID.String())
		if find == nil || find.Status == item.Status {
			continue
		}

		find.SetStatus(item.Status)
		if err := x.cloudRepo.Update(ctx, find); err != nil {
			multilerr = multierror.Append(multilerr, err)
			return &exceptions.CustomError{
				Status: exceptions.ERRSYSTEM,
				Errors: multilerr,
			}
		}
	}

	return nil
}

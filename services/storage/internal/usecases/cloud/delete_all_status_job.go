package cloud

import (
	"context"
	"log"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *cloudInteractor) DeleteAllStatusJob(ctx context.Context) *exceptions.CustomError {
	var multilerr *multierror.Error

	log.Println("start delete all status job")

	params := entities.CloudQueryParams{
		Status:    entities.CloudStatusPending,
		CreatedAt: time.Now().AddDate(0, 0, -1),
	}

	all, err := x.cloudRepo.All(ctx, &params)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRSYSTEM,
			Errors: multilerr,
		}
	}

	for _, find := range all {
		if err := x.cloudRepo.Delete(ctx, find.ID.String()); err != nil {
			multilerr = multierror.Append(multilerr, err)
			return &exceptions.CustomError{
				Status: exceptions.ERRSYSTEM,
				Errors: multilerr,
			}
		}
	}

	log.Println("success delete all status job")

	return nil
}

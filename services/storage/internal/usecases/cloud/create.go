package cloud

import (
	"context"
	"os"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *cloudInteractor) Create(ctx context.Context, payload entities.CloudDto) (*entities.Cloud, *exceptions.CustomError) {
	var multilerr *multierror.Error

	cloud := entities.NewCloud(payload)

	if payload.Url == "" {
		path, errUpload := x.awsService.UploadFile(ctx, &payload.File)
		if errUpload != nil {
			multilerr = multierror.Append(multilerr, errUpload)
			return nil, &exceptions.CustomError{
				Status: exceptions.ERRBUSSINESS,
				Errors: multilerr,
			}
		}

		cloud.SetUrl(path)
		cloud.SetStatus(entities.CloudStatusPending)

		if err := os.Remove(payload.File.Name); err != nil {
			multilerr = multierror.Append(multilerr, err)
			return nil, &exceptions.CustomError{
				Status: exceptions.ERRSYSTEM,
				Errors: multilerr,
			}
		}
	}

	if err := cloud.Validate(); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRDOMAIN,
			Errors: multilerr,
		}
	}

	if err := x.cloudRepo.Create(ctx, cloud); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return cloud, nil
}

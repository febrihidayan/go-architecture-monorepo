package device_token

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *deviceTokenInteractor) Create(ctx context.Context, payload entities.DeviceTokenDto) (*entities.DeviceToken, *exceptions.CustomError) {
	var multilerr *multierror.Error

	deviceToken := entities.NewDeviceToken(payload)
	if err := deviceToken.Validate(); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRDOMAIN,
			Errors: multilerr,
		}
	}

	if err := x.deviceTokenRepo.Create(ctx, deviceToken); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return deviceToken, nil
}

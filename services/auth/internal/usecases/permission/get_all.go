package permission

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *permissionInteractor) GetAll(ctx context.Context, params entities.PermissionQueryParams) (*entities.PermissionMeta, *exceptions.CustomError) {
	var multilerr *multierror.Error

	all, total, err := x.permissionRepo.GetAll(ctx, &params)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return &entities.PermissionMeta{
		Data:  all,
		Total: total,
	}, nil
}

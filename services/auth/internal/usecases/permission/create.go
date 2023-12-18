package permission

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *permissionInteractor) Create(ctx context.Context, payload entities.PermissionDto) (*entities.Permission, *exceptions.CustomError) {
	var multilerr *multierror.Error

	find, _ := x.permissionRepo.FindByName(ctx, payload.Name)
	if find != nil {
		multilerr = multierror.Append(multilerr, lang.ErrPermissionAlready)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	permission := entities.NewPermission(payload)
	if err := permission.Validate(); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRDOMAIN,
			Errors: multilerr,
		}
	}

	if err := x.permissionRepo.Create(ctx, permission); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return permission, nil
}

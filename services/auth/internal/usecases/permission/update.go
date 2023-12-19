package permission

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *permissionInteractor) Update(ctx context.Context, payload entities.PermissionDto) (*entities.Permission, *exceptions.CustomError) {
	var multilerr *multierror.Error

	find, err := x.permissionRepo.Find(ctx, payload.ID.String())
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	if find.Name != payload.Name {
		FindName, _ := x.permissionRepo.FindByName(ctx, payload.Name)
		if FindName != nil {
			multilerr = multierror.Append(multilerr, lang.ErrPermissionAlready)
			return nil, &exceptions.CustomError{
				Status: exceptions.ERRREPOSITORY,
				Errors: multilerr,
			}
		}
	}

	permission := entities.NewPermission(payload, find)
	if err := permission.Validate(); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRDOMAIN,
			Errors: multilerr,
		}
	}

	if err := x.permissionRepo.Update(ctx, permission); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return permission, nil
}

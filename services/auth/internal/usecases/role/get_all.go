package role

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *roleInteractor) GetAll(ctx context.Context, params entities.RoleQueryParams) (*entities.RoleMeta, *exceptions.CustomError) {
	var multilerr *multierror.Error

	all, total, err := x.roleRepo.GetAll(ctx, &params)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return &entities.RoleMeta{
		Data:  all,
		Total: total,
	}, nil
}

package role

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *roleInteractor) Update(ctx context.Context, payload entities.RoleDto) (*entities.Role, *exceptions.CustomError) {
	var multilerr *multierror.Error

	find, err := x.roleRepo.Find(ctx, payload.ID.String())
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	if find.Name != payload.Name {
		FindName, _ := x.roleRepo.FindByName(ctx, payload.Name)
		if FindName != nil {
			multilerr = multierror.Append(multilerr, lang.ErrRoleAlready)
			return nil, &exceptions.CustomError{
				Status: exceptions.ERRREPOSITORY,
				Errors: multilerr,
			}
		}
	}

	role := entities.NewRole(payload, find)
	if err := role.Validate(); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRDOMAIN,
			Errors: multilerr,
		}
	}

	if err := x.roleRepo.Update(ctx, role); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return role, nil
}

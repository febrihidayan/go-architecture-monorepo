package role

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *roleInteractor) Create(ctx context.Context, payload entities.RoleDto) (*entities.Role, *exceptions.CustomError) {
	var multilerr *multierror.Error

	find, _ := x.roleRepo.FindByName(ctx, payload.Name)
	if find != nil {
		multilerr = multierror.Append(multilerr, lang.ErrRoleAlready)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	role := entities.NewRole(payload)
	if err := role.Validate(); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRDOMAIN,
			Errors: multilerr,
		}
	}

	if err := x.roleRepo.Create(ctx, role); err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	return role, nil
}

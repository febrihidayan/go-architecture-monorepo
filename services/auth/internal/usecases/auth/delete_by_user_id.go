package auth

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"

	"github.com/hashicorp/go-multierror"
)

func (x *authInteractor) DeleteByUserID(ctx context.Context, userId string) *exceptions.CustomError {
	var (
		multilerr *multierror.Error
	)

	log.Println("DeleteByUserID::info#1", "start check userId already")
	auth, err := x.authRepo.FindByUserId(ctx, userId)
	if err != nil {
		log.Println("DeleteByUserID::error#1", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	log.Println("DeleteByUserID::info#2", "start deleting an existing user")
	if err := x.authRepo.DeleteByUserID(ctx, auth.UserId); err != nil {
		log.Println("DeleteByUserID::error#2", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	log.Println("DeleteByUserID::info#3", "start deleting an existing user role")
	if err := x.roleUserRepo.DeleteByUserId(ctx, auth.UserId); err != nil {
		log.Println("DeleteByUserID::error#3", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	log.Println("DeleteByUserID::info#4", "start deleting an existing user permission")
	if err := x.permissionUserRepo.DeleteByUserId(ctx, auth.UserId); err != nil {
		log.Println("DeleteByUserID::error#4", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRBUSSINESS,
			Errors: multilerr,
		}
	}

	log.Println("DeleteByUserID::success#1", "deleted successfully")

	return nil
}

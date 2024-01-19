package acl_handler

import (
	"context"
	"net/http"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/response"
)

func (x *aclHttpHandler) GetAllPermission(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.Background()
	)

	results, err := x.aclUsecase.GetAllPermission(ctx)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapPermissionListResponses(results))
}

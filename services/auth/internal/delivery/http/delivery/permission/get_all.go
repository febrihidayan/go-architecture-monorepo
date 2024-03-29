package permission_handler

import (
	"context"
	"net/http"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/request"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/response"
)

func (x *PermissionHttpHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = context.Background()
		query request.PermissionQueryParams
	)

	if err := utils.MapQueryParams(r, &query); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{err})
		return
	}

	params := entities.PermissionQueryParams{
		Search:  query.Search,
		Page:    query.Page,
		PerPage: query.PerPage,
	}

	results, err := x.PermissionUsecase.GetAll(ctx, params)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.ResponseWithJsonMeta(w, http.StatusOK, response.MapPermissionListResponses(results.Data), utils.MetaResponse{
		Total:   results.Total,
		Page:    params.Page,
		PerPage: params.PerPage,
	})
}

package notification_handler

import (
	"context"
	"net/http"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/http/request"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/http/response"
)

func (x *notificationHttpHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = context.Background()
		query request.NotificationQueryParams
	)

	if err := utils.MapQueryParams(r, &query); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{err})
		return
	}

	jwtToken, errJwt := utils.DecodeJwtToken(r.Header.Get("Authorization"))
	if errJwt != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errJwt})
		return
	}

	params := entities.NotificationQueryParams{
		UserId:  jwtToken.Subject,
		Page:    query.Page,
		PerPage: query.PerPage,
	}

	results, err := x.notificationUsecase.GetAll(ctx, params)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.ResponseWithJsonMeta(w, http.StatusOK, response.MapNotificationListResponses(results.Data), utils.MetaResponse{
		Total:   results.Total,
		Page:    results.Page,
		PerPage: results.PerPage,
	})
}

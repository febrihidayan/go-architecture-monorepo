package device_token_handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/validator"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/http/request"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/http/response"
)

func (x *DeviceTokenHttpHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.Background()
		payload request.DeviceTokenCreateRequest
	)

	jwtToken, errJwt := utils.DecodeJwtToken(r.Header.Get("Authorization"))
	if errJwt != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errJwt})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{err})
		return
	}

	if err := validator.Make(payload); err != nil {
		validator.ErrorJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	data := entities.DeviceTokenDto{
		UserId: jwtToken.Subject,
		Token:  payload.Token,
		OsName: payload.OsName,
	}

	result, err := x.DeviceTokenUsecase.Create(ctx, data)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, response.MapDeviceTokenListResponse(result))
}

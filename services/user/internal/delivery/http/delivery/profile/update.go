package profile_handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/validator"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/http/request"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/http/response"
)

func (x *ProfileHttpHandler) Update(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.Background()
		payload request.UserUpdateRequest
	)

	jwtToken, errJwt := utils.DecodeJwtToken(r.Header.Get("Authorization"))
	if errJwt != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errJwt})
		return
	}

	id, _ := common.StringToID(jwtToken.Subject)
	payload.ID = &id

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{err})
		return
	}

	if err := validator.Make(payload); err != nil {
		validator.ErrorJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	data := entities.UserDto{
		ID:       payload.ID,
		Name:     payload.Name,
		Email:    payload.Email,
		Avatar:   payload.Avatar,
		LangCode: payload.LangCode,
	}

	result, err := x.ProfileUsecase.Update(ctx, data)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapUserListResponse(result))
}

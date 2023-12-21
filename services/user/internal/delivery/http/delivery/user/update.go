package user_handler

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
	"github.com/gorilla/mux"
)

func (x *userHttpHandler) Update(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.Background()
		payload request.UserUpdateRequest
		vars    = mux.Vars(r)
		id, _   = common.StringToID(vars["id"])
	)

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
		ID:    payload.ID,
		Name:  payload.Name,
		Email: payload.Email,
	}

	result, err := x.userUsecase.Update(ctx, data)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapUserListResponse(result))
}

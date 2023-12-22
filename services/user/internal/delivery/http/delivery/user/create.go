package user_handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/validator"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/http/request"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/http/response"
)

func (x *userHttpHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.Background()
		payload request.UserCreateRequest
	)

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{err})
		return
	}

	if err := validator.Make(payload); err != nil {
		validator.ErrorJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	data := entities.UserDto{
		Name:  payload.Name,
		Email: payload.Email,
		Auth: entities.Auth{
			Password: payload.Password,
		},
	}

	result, err := x.userUsecase.Create(ctx, data)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, response.MapUserListResponse(result))
}

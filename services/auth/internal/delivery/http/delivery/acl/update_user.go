package acl_handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/validator"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/request"
	"github.com/gorilla/mux"
)

func (x *aclHttpHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.Background()
		vars    = mux.Vars(r)
		id      = vars["id"]
		payload request.AclUserUpdateRequest
	)

	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errors.New("param id required")})
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{err})
		return
	}

	if err := validator.Make(payload); err != nil {
		validator.ErrorJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	data := entities.AclUserDto{
		UserId:      id,
		Permissions: payload.Permissions,
		Roles:       payload.Roles,
	}

	if err := x.aclUsecase.UpdateUser(ctx, data); err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, nil)
}

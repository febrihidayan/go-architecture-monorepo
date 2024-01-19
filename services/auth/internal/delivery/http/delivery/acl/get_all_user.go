package acl_handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/response"
	"github.com/gorilla/mux"
)

func (x *aclHttpHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.Background()
		vars = mux.Vars(r)
		id   = vars["id"]
	)

	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errors.New("param id required")})
		return
	}

	_, errJwt := utils.DecodeJwtToken(r.Header.Get("Authorization"))
	if errJwt != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errJwt})
		return
	}

	results, err := x.aclUsecase.GetAllUser(ctx, id)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response.MapAclListResponse(results))
}

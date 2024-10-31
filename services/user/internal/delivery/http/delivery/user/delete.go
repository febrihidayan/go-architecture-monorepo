package user_handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/gorilla/mux"
)

func (x *UserHttpHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.Background()
		vars = mux.Vars(r)
		id   = vars["id"]
	)

	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errors.New("param id required")})
		return
	}

	if err := x.UserUsecase.Delete(ctx, id); err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, nil)
}

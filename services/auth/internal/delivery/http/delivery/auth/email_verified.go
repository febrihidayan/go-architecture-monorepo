package auth_handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/gorilla/mux"
)

func (x *authHttpHandler) EmailVerified(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = context.Background()
		vars  = mux.Vars(r)
		token = vars["token"]
	)

	if token == "" {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errors.New("param token required")})
		return
	}

	if err := x.authUsecase.EmailVerified(ctx, token); err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, nil)
}

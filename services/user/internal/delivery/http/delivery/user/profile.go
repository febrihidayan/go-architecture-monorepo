package user_handler

import (
	"context"
	"net/http"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
)

func (x *userHttpHandler) Profile(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.Background()
	)

	_, err := x.userUsecase.Profile(ctx, "")
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, nil)
}

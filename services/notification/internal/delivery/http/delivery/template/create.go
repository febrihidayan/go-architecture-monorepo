package template_handler

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

func (x *TemplateHttpHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.Background()
		payload request.TemplateCreateRequest
	)

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{err})
		return
	}

	if err := validator.Make(payload); err != nil {
		validator.ErrorJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	dataJson, errJson := json.Marshal(payload.Data)
	if errJson != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errJson})
		return
	}

	data := entities.TemplateDto{
		Name: payload.Name,
		Data: string(dataJson),
	}

	result, err := x.TemplateUsecase.Create(ctx, data)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(err.Status), err.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, response.MapTemplateListResponse(result))
}

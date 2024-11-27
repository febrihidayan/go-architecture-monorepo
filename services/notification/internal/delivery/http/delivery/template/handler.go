package template_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/template"

	"github.com/gorilla/mux"
)

type TemplateHttpHandler struct {
	Cfg             *config.NotificationConfig
	TemplateUsecase usecases.TemplateUsecase
}

func NewTemplateHttpHandler(
	r *mux.Router,
	deps *factories.Dependencies,
) {
	handler := &TemplateHttpHandler{
		Cfg:             deps.Config,
		TemplateUsecase: template.NewTemplateInteractor(deps),
	}

	r.HandleFunc("/v1/notification/template", handler.Create).Methods("POST")
}

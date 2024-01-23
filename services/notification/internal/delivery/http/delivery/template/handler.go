package template_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/template"

	"github.com/gorilla/mux"
)

type TemplateHttpHandler struct {
	Cfg             *config.NotificationConfig
	TemplateUsecase usecases.TemplateUsecase
}

func NewTemplateHttpHandler(
	r *mux.Router,
	config *config.NotificationConfig,
	mongoFactory *factories.MongoFactory,
) {
	handler := &TemplateHttpHandler{
		Cfg: config,
		TemplateUsecase: template.NewTemplateInteractor(
			config,
			mongoFactory,
		),
	}

	r.HandleFunc("/v1/notification/template", handler.Create).Methods("POST")
}

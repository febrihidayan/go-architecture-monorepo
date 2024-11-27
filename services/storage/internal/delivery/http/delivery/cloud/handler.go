package cloud_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/usecases/cloud"

	"github.com/gorilla/mux"
)

type CloudHttpHandler struct {
	Cfg          *config.StorageConfig
	CloudUsecase usecases.CloudUsecase
}

func NewCloudHttpHandler(
	r *mux.Router,
	deps *factories.Dependencies,
) {
	handler := &CloudHttpHandler{
		Cfg:          deps.Config,
		CloudUsecase: cloud.NewCloudInteractor(deps),
	}

	r.HandleFunc("/v1/storage/cloud", handler.Create).Methods("POST")
}

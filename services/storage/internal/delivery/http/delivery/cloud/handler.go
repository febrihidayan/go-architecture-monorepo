package cloud_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/services"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/usecases/cloud"

	"github.com/gorilla/mux"
)

type CloudHttpHandler struct {
	Cfg          *config.StorageConfig
	CloudUsecase usecases.CloudUsecase
}

func NewCloudHttpHandler(
	r *mux.Router,
	config *config.StorageConfig,
	mongoFactory *factories.MongoFactory,
	awsService services.AwsService,
) {
	handler := &CloudHttpHandler{
		Cfg: config,
		CloudUsecase: cloud.NewCloudInteractor(
			config,
			mongoFactory,
			&awsService,
		),
	}

	r.HandleFunc("/v1/storage/cloud", handler.Create).Methods("POST")
}

package cloud_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/services"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/usecases/cloud"

	"github.com/gorilla/mux"
)

type cloudHttpHandler struct {
	cfg          *config.StorageConfig
	cloudUsecase usecases.CloudUsecase
}

func CloudHttpHandler(
	r *mux.Router,
	config *config.StorageConfig,
	mongoFactory *factories.MongoFactory,
	awsService services.AwsService,
) {
	handler := &cloudHttpHandler{
		cfg: config,
		cloudUsecase: cloud.NewCloudInteractor(
			config,
			mongoFactory,
			&awsService,
		),
	}

	r.HandleFunc("/v1/storage/cloud", handler.Create).Methods("POST")
}

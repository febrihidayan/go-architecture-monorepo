package cloud_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	repository_mongo "github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/mongo"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/services"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/usecases/cloud"

	"github.com/gorilla/mux"
)

type cloudHttpHandler struct {
	cfg          *config.StorageConfig
	cloudUsecase usecases.CloudUsecase
}

func RoleHttpHandler(
	r *mux.Router,
	config *config.StorageConfig,
	cloudRepo repository_mongo.CloudRepository,
	awsService services.AwsService,
) {
	handler := &cloudHttpHandler{
		cfg: config,
		cloudUsecase: cloud.NewCloudInteractor(
			config,
			&cloudRepo,
			&awsService,
		),
	}

	r.HandleFunc("/v1/storage/cloud", handler.Create).Methods("POST")
}

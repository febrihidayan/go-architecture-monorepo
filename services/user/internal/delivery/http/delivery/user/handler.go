package user_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	repository_grpc "github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/grpc"
	repository_mongo "github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/mongo"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/usecases/user"

	"github.com/gorilla/mux"
)

type userHttpHandler struct {
	cfg         *config.UserConfig
	userUsecase usecases.UserUsecase
}

func UserHttpHandler(
	r *mux.Router,
	config *config.UserConfig,
	userRepo repository_mongo.UserRepository,
	authRepo repository_grpc.AuthRepository,
) {
	handler := &userHttpHandler{
		cfg: config,
		userUsecase: user.NewUserInteractor(
			config,
			&userRepo,
			&authRepo,
		),
	}

	r.HandleFunc("/v1/users", handler.GetAll).Methods("GET")
	r.HandleFunc("/v1/user", handler.Create).Methods("POST")
	r.HandleFunc("/v1/user/{id}", handler.Update).Methods("PUT")
}

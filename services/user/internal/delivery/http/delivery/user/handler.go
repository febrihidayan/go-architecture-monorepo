package user_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/rabbitmq_server/publisher"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/usecases/user"

	"github.com/gorilla/mux"
)

type UserHttpHandler struct {
	Cfg         *config.UserConfig
	UserUsecase usecases.UserUsecase
}

func NewUserHttpHandler(
	r *mux.Router,
	config *config.UserConfig,
	mongoFactory *factories.MongoFactory,
	grpcFactory *factories.GrpcClientFactory,
	rabbitmq *publisher.PublisherRabbitMQ) {
	handler := &UserHttpHandler{
		Cfg: config,
		UserUsecase: user.NewUserInteractor(
			config,
			mongoFactory,
			grpcFactory,
			rabbitmq,
		),
	}

	r.HandleFunc("/v1/users", handler.GetAll).Methods("GET")
	r.HandleFunc("/v1/user", handler.Create).Methods("POST")
	r.HandleFunc("/v1/user/{id}", handler.Find).Methods("GET")
	r.HandleFunc("/v1/user/{id}", handler.Update).Methods("PUT")
	r.HandleFunc("/v1/user/{id}", handler.Delete).Methods("DELETE")
}

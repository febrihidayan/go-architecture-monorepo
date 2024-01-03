package grpc_server

import (
	userPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/user"

	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/grpc_client"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/usecases/profile"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/usecases/user"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func HandlerUserServices(s *grpc.Server, db *mongo.Database, cfg config.UserConfig) {

	mongoFactory := factories.NewMongoFactory(db)

	grpcClient, _ := grpc_client.NewGrpcClient(&cfg.GrpcClient)
	grpcClientFactory := factories.NewGrpcFactory(grpcClient)

	userPb.RegisterUserServicesServer(s, &server{
		userUsecase:    user.NewUserInteractor(&cfg, mongoFactory, grpcClientFactory),
		profileUsecase: profile.NewProfileInteractor(&cfg, mongoFactory, grpcClientFactory),
	})
}

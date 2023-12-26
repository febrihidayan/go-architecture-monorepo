package grpc_server

import (
	authPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/auth"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/grpc_client"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func HandlerAuthServices(s *grpc.Server, db *mongo.Database, cfg config.AuthConfig) {
	mongoFactory := factories.NewMongoFactory(db)

	grpcClient, _ := grpc_client.NewGrpcClient(&cfg.GrpcClient)
	grpcClientFactory := factories.NewGrpcFactory(grpcClient)

	authPb.RegisterAuthServicesServer(s, &server{
		authUsecase: auth.NewAuthInteractor(&cfg, mongoFactory, grpcClientFactory),
	})
}

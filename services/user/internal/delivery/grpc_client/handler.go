package grpc_client

import (
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	"google.golang.org/grpc"
)

func NewGrpcClient(cfg *config.GrpcClient) (*ServerClient, []error) {
	var (
		client ServerClient
		err    error
		errs   []error
	)

	client.AuthClient, err = grpc.Dial(cfg.Auth, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect grpc auth: %v", err)

		errs = append(errs, err)
	}
	log.Println("rpc auth started on", cfg.Auth)

	client.StorageClient, err = grpc.Dial(cfg.Storage, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect grpc storage: %v", err)

		errs = append(errs, err)
	}
	log.Println("rpc storage started on", cfg.Storage)

	return &client, errs
}

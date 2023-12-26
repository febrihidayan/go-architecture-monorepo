package grpc_client

import (
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"google.golang.org/grpc"
)

func NewGrpcClient(cfg *config.GrpcClient) (*ServerClient, []error) {
	var (
		client ServerClient
		err    error
		errs   []error
	)

	client.UserClient, err = grpc.Dial(cfg.User, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect grpc user: %v", err)

		errs = append(errs, err)
	}
	log.Println("rpc user started on", cfg.User)

	return &client, errs
}

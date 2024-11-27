package factories

import (
	"fmt"
	"log"
	"sync"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/grpc_client"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dependencies struct {
	Config            *config.AuthConfig
	MongoDB           *mongo.Database
	GrpcClient        *grpc_client.ServerClient
	RabbitMQConn      *rabbitmq.RabbitMQ
	MongoFactory      *MongoFactory
	GrpcClientFactory *GrpcClientFactory
}

var (
	dependencies *Dependencies
	once         sync.Once
)

func InitializeDependencies() *Dependencies {
	cfg := config.Auth()

	once.Do(func() {
		log.Println("Initializing Dependencies...")

		// MongoDB
		db := config.InitDatabaseMongodb()
		log.Println("MongoDB Initialized")

		// gRPC Client
		grpcClient, errs := grpc_client.NewGrpcClient(&cfg.GrpcClient)
		if len(errs) > 0 {
			log.Fatalf("Failed to initialize gRPC client: %v", errs)
		}

		// RabbitMQ
		dns := fmt.Sprintf(
			"amqp://%s:%s@%s:%s/",
			cfg.RabbitMQ.User,
			cfg.RabbitMQ.Password,
			cfg.RabbitMQ.Host,
			cfg.RabbitMQ.Port,
		)
		rmq, err := rabbitmq.NewRabbitMQ(dns)
		if err != nil {
			log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		}

		dependencies = &Dependencies{
			Config:            cfg,
			MongoDB:           db,
			GrpcClient:        grpcClient,
			RabbitMQConn:      rmq,
			MongoFactory:      NewMongoFactory(db),
			GrpcClientFactory: NewGrpcFactory(grpcClient),
		}
	})

	return dependencies
}

func GetDependencies() *Dependencies {
	if dependencies == nil {
		log.Fatal("Dependencies not initialized")
	}
	return dependencies
}

func (x *Dependencies) Close() {
	log.Println("Closing dependencies...")
	if err := x.MongoDB.Client().Disconnect(nil); err != nil {
		log.Printf("Failed to disconnect MongoDB: %v", err)
	}

	x.RabbitMQConn.Close()

	// No explicit close for gRPC Client
	log.Println("Dependencies closed")
}

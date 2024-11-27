package factories

import (
	"log"
	"sync"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/grpc_client"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dependencies struct {
	Config            *config.NotificationConfig
	MongoDB           *mongo.Database
	GrpcClient        *grpc_client.ServerClient
	MongoFactory      *MongoFactory
	GrpcClientFactory *GrpcClientFactory
	FirebaseGoogle    *services.FirebaseGoogleService
	Mailgun           *services.MailgunService
}

var (
	dependencies *Dependencies
	once         sync.Once
)

func InitializeDependencies() *Dependencies {
	cfg := config.Notification()

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

		// run firebase google service
		firebaseGoogleService, errFGService := services.NewFcmGoogleService(cfg.FirebaseGoogleService)
		if errFGService != nil {
			log.Fatalf("did not connect firebase service: %v", errFGService)
		}

		// run mailgun service
		mailgun := services.NewMailgunClient(cfg.Mailgun)

		dependencies = &Dependencies{
			Config:            cfg,
			MongoDB:           db,
			GrpcClient:        grpcClient,
			FirebaseGoogle:    firebaseGoogleService,
			Mailgun:           mailgun,
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

	// No explicit close for gRPC Client
	log.Println("Dependencies closed")
}

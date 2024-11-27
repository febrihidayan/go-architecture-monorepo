package factories

import (
	"fmt"
	"log"
	"sync"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dependencies struct {
	Config       *config.StorageConfig
	MongoDB      *mongo.Database
	RabbitMQConn *rabbitmq.RabbitMQ
	MongoFactory *MongoFactory
	AwsService   *services.AwsService
}

var (
	dependencies *Dependencies
	once         sync.Once
)

func InitializeDependencies() *Dependencies {
	cfg := config.Storage()

	once.Do(func() {
		log.Println("Initializing Dependencies...")

		// MongoDB
		db := config.InitDatabaseMongodb()
		log.Println("MongoDB Initialized")

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

		// run AWS Service
		awsService := services.NewAwsService(&cfg.Aws)

		dependencies = &Dependencies{
			Config:       cfg,
			MongoDB:      db,
			RabbitMQConn: rmq,
			AwsService:   awsService,
			MongoFactory: NewMongoFactory(db),
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

package config

import (
	"context"
	"fmt"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitDatabaseMongodb() *mongo.Database {
	cfg := config.DatabaseMongodb()

	dns := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)

	client, err := mongo.NewClient(options.Client().ApplyURI(dns))
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.TODO())

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		panic(err)
	}

	return client.Database(cfg.Database)
}

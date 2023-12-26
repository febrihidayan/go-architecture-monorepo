package factories

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/repositories"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoFactory struct {
	CloudRepo repositories.CloudRepository
}

func NewMongoFactory(db *mongo.Database) *MongoFactory {
	var (
		CloudRepo = mongo_repositories.NewCloudRepository(db)
	)

	return &MongoFactory{
		CloudRepo: &CloudRepo,
	}
}

package factories

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/repositories"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoFactory struct {
	UserRepo repositories.UserRepository
}

func NewMongoFactory(db *mongo.Database) *MongoFactory {
	var (
		UserRepo = mongo_repositories.NewUserRepository(db)
	)

	return &MongoFactory{
		UserRepo: &UserRepo,
	}
}

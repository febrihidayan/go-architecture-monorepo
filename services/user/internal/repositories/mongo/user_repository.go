package repositories

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/mongo/mappers"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return UserRepository{db: db}
}

func (x *UserRepository) Create(ctx context.Context, payload *entities.User) error {
	_, err := x.db.Collection(models.User{}.TableName()).InsertOne(ctx, mappers.ToModelUser(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *UserRepository) Find(ctx context.Context, id string) (*entities.User, error) {
	var auth models.User

	err := x.db.Collection(models.User{}.TableName()).FindOne(ctx, bson.M{"_id": id}).Decode(&auth)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return mappers.ToDomainUser(&auth), nil
}

func (x *UserRepository) Update(ctx context.Context, payload *entities.User) error {
	_, err := x.db.Collection(models.User{}.TableName()).ReplaceOne(ctx, bson.M{
		"_id": payload.ID.String(),
	}, mappers.ToModelUser(payload))

	if err != nil {
		return err
	}

	return nil
}

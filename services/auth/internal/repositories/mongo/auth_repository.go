package repositories

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/mongoqb"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/mappers"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	db *mongoqb.MongoQueryBuilder
}

func NewAuthRepository(db *mongo.Database) AuthRepository {
	return AuthRepository{
		db: mongoqb.NewMongoQueryBuilder(db.Collection(models.Auth{}.TableName())),
	}
}

func (x *AuthRepository) Create(ctx context.Context, payload *entities.Auth) error {
	_, err := x.db.InsertOne(ctx, mappers.ToModelAuth(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *AuthRepository) Find(ctx context.Context, id string) (*entities.Auth, error) {
	var auth models.Auth

	err := x.db.FindByID(ctx, id).Decode(&auth)

	if err != nil {
		return nil, errors.New("auth not found")
	}

	return mappers.ToDomainAuth(&auth), nil
}

func (x *AuthRepository) FindByEmail(ctx context.Context, email string) (*entities.Auth, error) {
	var auth models.Auth

	err := x.db.FindOne(ctx, bson.M{"email": email}).Decode(&auth)

	if err != nil {
		return nil, errors.New("auth not found")
	}

	return mappers.ToDomainAuth(&auth), nil
}

func (x *AuthRepository) FindByUserId(ctx context.Context, userId string) (*entities.Auth, error) {
	var auth models.Auth

	err := x.db.FindOne(ctx, bson.M{"user_id": userId}).Decode(&auth)

	if err != nil {
		return nil, errors.New("auth not found")
	}

	return mappers.ToDomainAuth(&auth), nil
}

func (x *AuthRepository) Update(ctx context.Context, payload *entities.Auth) error {
	_, err := x.db.ReplaceOne(ctx, bson.M{
		"_id": payload.ID.String(),
	}, mappers.ToModelAuth(payload))

	if err != nil {
		return err
	}

	return nil
}

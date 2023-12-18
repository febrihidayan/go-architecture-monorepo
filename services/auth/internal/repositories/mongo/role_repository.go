package repositories

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/mappers"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoleRepository struct {
	db *mongo.Database
}

func NewRoleRepository(db *mongo.Database) RoleRepository {
	return RoleRepository{db: db}
}

func (x *RoleRepository) Create(ctx context.Context, payload *entities.Role) error {
	_, err := x.db.Collection(models.Role{}.TableName()).InsertOne(ctx, mappers.ToModelRole(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *RoleRepository) Find(ctx context.Context, id string) (*entities.Role, error) {
	var role models.Role

	err := x.db.Collection(models.Role{}.TableName()).FindOne(ctx, bson.M{"_id": id}).Decode(&role)

	if err != nil {
		return nil, errors.New("role not found")
	}

	return mappers.ToDomainRole(&role), nil
}

func (x *RoleRepository) FindByName(ctx context.Context, name string) (*entities.Role, error) {
	var role models.Role

	err := x.db.Collection(models.Role{}.TableName()).FindOne(ctx, bson.M{"name": name}).Decode(&role)

	if err != nil {
		return nil, errors.New("role not found")
	}

	return mappers.ToDomainRole(&role), nil
}

func (x *RoleRepository) Update(ctx context.Context, payload *entities.Role) error {
	_, err := x.db.Collection(models.Role{}.TableName()).ReplaceOne(ctx, bson.M{
		"_id": payload.ID.String(),
	}, mappers.ToModelRole(payload))

	if err != nil {
		return err
	}

	return nil
}

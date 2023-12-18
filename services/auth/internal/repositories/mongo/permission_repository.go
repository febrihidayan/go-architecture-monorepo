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

type PermissionRepository struct {
	db *mongo.Database
}

func NewPermissionRepository(db *mongo.Database) PermissionRepository {
	return PermissionRepository{db: db}
}

func (x *PermissionRepository) Create(ctx context.Context, payload *entities.Permission) error {
	_, err := x.db.Collection(models.Permission{}.TableName()).InsertOne(ctx, mappers.ToModelPermission(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *PermissionRepository) Find(ctx context.Context, id string) (*entities.Permission, error) {
	var permission models.Permission

	err := x.db.Collection(models.Permission{}.TableName()).FindOne(ctx, bson.M{"_id": id}).Decode(&permission)

	if err != nil {
		return nil, errors.New("permission not found")
	}

	return mappers.ToDomainPermission(&permission), nil
}

func (x *PermissionRepository) FindByName(ctx context.Context, name string) (*entities.Permission, error) {
	var permission models.Permission

	err := x.db.Collection(models.Permission{}.TableName()).FindOne(ctx, bson.M{"name": name}).Decode(&permission)

	if err != nil {
		return nil, errors.New("permission not found")
	}

	return mappers.ToDomainPermission(&permission), nil
}

func (x *PermissionRepository) Update(ctx context.Context, payload *entities.Permission) error {
	_, err := x.db.Collection(models.Permission{}.TableName()).ReplaceOne(ctx, bson.M{
		"_id": payload.ID.String(),
	}, mappers.ToModelPermission(payload))

	if err != nil {
		return err
	}

	return nil
}

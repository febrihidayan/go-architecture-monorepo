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

type PermissionUserRepository struct {
	db *mongo.Database
}

func NewPermissionUserRepository(db *mongo.Database) PermissionUserRepository {
	return PermissionUserRepository{db: db}
}

func (x *PermissionUserRepository) CreateMany(ctx context.Context, payloads []*entities.PermissionUser) error {
	_, err := x.db.Collection(models.PermissionUser{}.TableName()).InsertMany(ctx, mappers.ToListModelPermissionUser(payloads))

	if err != nil {
		return err
	}

	return nil
}

func (x *PermissionUserRepository) AllByUserId(ctx context.Context, userId string) ([]*entities.PermissionUser, error) {
	var roles []*models.PermissionUser

	cursor, err := x.db.Collection(models.PermissionUser{}.TableName()).Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &roles); err != nil {
		return nil, errors.New("role not found")
	}

	return mappers.ToListDomainPermissionUser(roles), nil
}

func (x *PermissionUserRepository) DeleteByPermissionIds(ctx context.Context, ids []string) error {
	_, err := x.db.Collection(models.PermissionUser{}.TableName()).DeleteMany(ctx, bson.M{
		"permission_id": bson.D{{"$in", ids}},
	})

	return err
}

func (x *PermissionUserRepository) DeleteByUserId(ctx context.Context, userId string) error {
	_, err := x.db.Collection(models.PermissionUser{}.TableName()).DeleteMany(ctx, bson.M{
		"user_id": userId,
	})

	return err
}

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

type RoleUserRepository struct {
	db *mongo.Database
}

func NewRoleUserRepository(db *mongo.Database) RoleUserRepository {
	return RoleUserRepository{db: db}
}

func (x *RoleUserRepository) CreateMany(ctx context.Context, payloads []*entities.RoleUser) error {
	_, err := x.db.Collection(models.RoleUser{}.TableName()).InsertMany(ctx, mappers.ToListModelRoleUser(payloads))

	if err != nil {
		return err
	}

	return nil
}

func (x *RoleUserRepository) AllByUserId(ctx context.Context, userId string) ([]*entities.RoleUser, error) {
	var roles []*models.RoleUser

	cursor, err := x.db.Collection(models.RoleUser{}.TableName()).Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &roles); err != nil {
		return nil, errors.New("role not found")
	}

	return mappers.ToListDomainRoleUser(roles), nil
}

func (x *RoleUserRepository) DeleteByRoleIds(ctx context.Context, ids []string) error {
	_, err := x.db.Collection(models.RoleUser{}.TableName()).DeleteMany(ctx, bson.M{
		"role_id": bson.D{{"$in", ids}},
	})

	return err
}

func (x *RoleUserRepository) DeleteByUserId(ctx context.Context, userId string) error {
	_, err := x.db.Collection(models.RoleUser{}.TableName()).DeleteMany(ctx, bson.M{
		"user_id": userId,
	})

	return err
}

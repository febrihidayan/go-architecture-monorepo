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

type PermissionRoleRepository struct {
	db *mongo.Database
}

func NewPermissionRoleRepository(db *mongo.Database) PermissionRoleRepository {
	return PermissionRoleRepository{db: db}
}

func (x *PermissionRoleRepository) CreateMany(ctx context.Context, payloads []*entities.PermissionRole) error {
	_, err := x.db.Collection(models.PermissionRole{}.TableName()).InsertMany(ctx, mappers.ToListModelPermissionRole(payloads))

	if err != nil {
		return err
	}

	return nil
}

func (x *PermissionRoleRepository) AllByRoleId(ctx context.Context, roleId string) ([]*entities.PermissionRole, error) {
	var roles []*models.PermissionRole

	cursor, err := x.db.Collection(models.PermissionRole{}.TableName()).Find(ctx, bson.M{"role_id": roleId})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &roles); err != nil {
		return nil, errors.New("permission role not found")
	}

	return mappers.ToListDomainPermissionRole(roles), nil
}

func (x *PermissionRoleRepository) Delete(ctx context.Context, payload *entities.PermissionRole) error {
	_, err := x.db.Collection(models.PermissionRole{}.TableName()).DeleteOne(ctx, bson.M{
		"role_id":       payload.RoleId,
		"permission_id": payload.PermissionId,
	})

	return err
}

func (x *PermissionRoleRepository) DeleteByPermissionIds(ctx context.Context, ids []string) error {
	_, err := x.db.Collection(models.PermissionRole{}.TableName()).DeleteMany(ctx, bson.M{
		"permission_id": bson.D{{"$in", ids}},
	})

	return err
}

func (x *PermissionRoleRepository) DeleteByRoleId(ctx context.Context, roleId string) error {
	_, err := x.db.Collection(models.PermissionRole{}.TableName()).DeleteMany(ctx, bson.M{
		"role_id": roleId,
	})

	return err
}

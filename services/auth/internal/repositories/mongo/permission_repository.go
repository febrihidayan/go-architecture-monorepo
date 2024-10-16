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

type PermissionRepository struct {
	db *mongoqb.MongoQueryBuilder
}

func NewPermissionRepository(db *mongo.Database) PermissionRepository {
	return PermissionRepository{
		db: mongoqb.NewMongoQueryBuilder(db.Collection(models.Permission{}.TableName())),
	}
}

func (x *PermissionRepository) Create(ctx context.Context, payload *entities.Permission) error {
	_, err := x.db.InsertOne(ctx, mappers.ToModelPermission(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *PermissionRepository) Find(ctx context.Context, id string) (*entities.Permission, error) {
	var permission models.Permission

	err := x.db.FindByID(ctx, id).Decode(&permission)

	if err != nil {
		return nil, errors.New("permission not found")
	}

	return mappers.ToDomainPermission(&permission), nil
}

func (x *PermissionRepository) FindByName(ctx context.Context, name string) (*entities.Permission, error) {
	var permission models.Permission

	err := x.db.FindOne(ctx, bson.M{"name": name}).Decode(&permission)

	if err != nil {
		return nil, errors.New("permission not found")
	}

	return mappers.ToDomainPermission(&permission), nil
}

func (x *PermissionRepository) All(ctx context.Context) ([]*entities.Permission, error) {
	var roles []*models.Permission

	cursor, err := x.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &roles); err != nil {
		return nil, errors.New("permissions not found")
	}

	return mappers.ToListDomainPermission(roles), nil
}

func (x *PermissionRepository) GetAll(ctx context.Context, params *entities.PermissionQueryParams) ([]*entities.Permission, int, error) {
	query := x.db.NewPipeline()

	if params.Search != "" {
		query.SearchSingleField("name", params.Search)
	}

	query.
		Sort("created_at", false).
		CountFacet().
		Unwind("$total").
		Paginate(params.Page, params.PerPage)

	cursor, err := query.Execute(ctx)
	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close(ctx)

	var result models.PermissionMeta
	for cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, result.Total, err
		}
	}

	return mappers.ToListDomainPermission(result.Data), result.Total, nil
}

func (x *PermissionRepository) AllByUserId(ctx context.Context, userId string) ([]*entities.Permission, error) {
	var (
		results []*models.Permission
		query   = x.db.NewPipeline()
	)

	query.
		Lookup("permission_role", "_id", "permission_id", "permission_role").
		Lookup("role_user", "permission_role.role_id", "role_id", "role_user", bson.A{
			bson.D{{
				"$match", bson.D{
					{"user_id", userId},
				},
			}},
		}).
		Lookup("permission_user", "_id", "permission_id", "permission_user", bson.A{
			bson.D{{
				"$match", bson.D{
					{"user_id", userId},
				},
			}},
		}).
		Match(bson.D{{
			"$or", bson.A{
				bson.D{{"permission_user.user_id", userId}},
				bson.D{{"role_user.user_id", userId}},
			},
		}})

	cursor, err := query.Execute(ctx)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &results); err != nil {
		return nil, errors.New("permissions not found")
	}

	return mappers.ToListDomainPermission(results), nil
}

func (x *PermissionRepository) Update(ctx context.Context, payload *entities.Permission) error {
	_, err := x.db.ReplaceOne(ctx, bson.M{
		"_id": payload.ID.String(),
	}, mappers.ToModelPermission(payload))

	if err != nil {
		return err
	}

	return nil
}

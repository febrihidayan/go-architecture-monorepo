package repositories

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/mappers"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (x *PermissionRepository) GetAll(ctx context.Context, params *entities.PermissionQueryParams) ([]*entities.Permission, int, error) {
	var (
		filter = mongo.Pipeline{}
		match  bson.D
		skip   = (params.Page - 1) * params.PerPage
	)

	if params.Search != "" {
		match = append(match, bson.D{{"name", primitive.Regex{
			Pattern: params.Search,
			Options: "i",
		}}}...)
	}

	if len(match) > 0 {
		filter = append(filter, mongo.Pipeline{
			bson.D{{
				"$match", match,
			}},
		}...)
	}

	filter = append(filter, mongo.Pipeline{
		bson.D{{
			"$sort", bson.D{
				{"created_at", -1},
			},
		}},
		bson.D{{
			"$facet", bson.D{
				{"total", bson.A{
					bson.D{{
						"$count", "count",
					}},
				}},
				{"data", bson.A{
					bson.D{{
						"$addFields", bson.D{
							{"_id", "$_id"},
						},
					}},
				}},
			},
		}},
		bson.D{{
			"$unwind", "$total",
		}},
		bson.D{{
			"$project", bson.D{
				{"data", bson.D{
					{"$slice", bson.A{
						"$data", skip, bson.D{
							{"$ifNull", bson.A{
								params.PerPage, "$total.count",
							}},
						},
					}},
				}},
				{"page", bson.D{
					{"$literal", skip/params.PerPage + 1},
				}},
				{"per_page", bson.D{
					{"$literal", params.PerPage},
				}},
				{"total", "$total.count"},
			},
		}},
	}...)

	cursor, err := x.db.Collection(models.Permission{}.TableName()).Aggregate(ctx, filter)
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

func (x *PermissionRepository) Update(ctx context.Context, payload *entities.Permission) error {
	_, err := x.db.Collection(models.Permission{}.TableName()).ReplaceOne(ctx, bson.M{
		"_id": payload.ID.String(),
	}, mappers.ToModelPermission(payload))

	if err != nil {
		return err
	}

	return nil
}

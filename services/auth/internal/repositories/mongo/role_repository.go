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

func (x *RoleRepository) GetAll(ctx context.Context, params *entities.RoleQueryParams) ([]*entities.Role, int, error) {
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

	cursor, err := x.db.Collection(models.Role{}.TableName()).Aggregate(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close(ctx)

	var result models.RoleMeta
	for cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, result.Total, err
		}
	}

	return mappers.ToListDomainRole(result.Data), result.Total, nil
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

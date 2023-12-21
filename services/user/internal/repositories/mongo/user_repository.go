package repositories

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/mongo/mappers"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var item models.User

	err := x.db.Collection(models.User{}.TableName()).FindOne(ctx, bson.M{"_id": id}).Decode(&item)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return mappers.ToDomainUser(&item), nil
}

func (x *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	var item models.User

	err := x.db.Collection(models.User{}.TableName()).FindOne(ctx, bson.M{"email": email}).Decode(&item)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return mappers.ToDomainUser(&item), nil
}

func (x *UserRepository) GetAll(ctx context.Context, params *entities.UserQueryParams) ([]*entities.User, int, error) {
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

	cursor, err := x.db.Collection(models.User{}.TableName()).Aggregate(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close(ctx)

	var result models.UserMeta
	for cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, result.Total, err
		}
	}

	return mappers.ToListDomainUser(result.Data), result.Total, nil
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

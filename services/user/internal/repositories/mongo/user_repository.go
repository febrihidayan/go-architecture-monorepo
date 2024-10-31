package repositories

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/mongoqb"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/mongo/mappers"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongoqb.MongoQueryBuilder
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return UserRepository{
		db: mongoqb.NewMongoQueryBuilder(db.Collection(models.User{}.TableName())),
	}
}

func (x *UserRepository) Create(ctx context.Context, payload *entities.User) error {
	_, err := x.db.InsertOne(ctx, mappers.ToModelUser(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *UserRepository) Find(ctx context.Context, id string) (*entities.User, error) {
	var item models.User

	err := x.db.FindOne(ctx, bson.M{"_id": id}).Decode(&item)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return mappers.ToDomainUser(&item), nil
}

func (x *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	var item models.User

	err := x.db.FindOne(ctx, bson.M{"email": email}).Decode(&item)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return mappers.ToDomainUser(&item), nil
}

func (x *UserRepository) GetAll(ctx context.Context, params *entities.UserQueryParams) ([]*entities.User, int, error) {
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

	var result models.UserMeta
	for cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, result.Total, err
		}
	}

	return mappers.ToListDomainUser(result.Data), result.Total, nil
}

func (x *UserRepository) Update(ctx context.Context, payload *entities.User) error {
	_, err := x.db.ReplaceOne(ctx, bson.M{
		"_id": payload.ID.String(),
	}, mappers.ToModelUser(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *UserRepository) Delete(ctx context.Context, id string) error {
	_, err := x.db.DeleteOne(ctx, bson.M{
		"_id": id,
	})

	if err != nil {
		return err
	}

	return nil
}

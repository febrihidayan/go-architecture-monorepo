package mongoqb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoQueryBuilder is a struct to build MongoDB queries and aggregation pipelines.
type MongoQueryBuilder struct {
	pipeline   mongo.Pipeline
	collection *mongo.Collection
}

// NewMongoQueryBuilder creates a new instance of MongoQueryBuilder with a given collection.
func NewMongoQueryBuilder(collection *mongo.Collection) *MongoQueryBuilder {
	return &MongoQueryBuilder{
		collection: collection,
	}
}

// AddStage appends a custom stage to the MongoDB aggregation pipeline.
func (x *MongoQueryBuilder) AddStage(stage bson.D) *MongoQueryBuilder {
	x.pipeline = append(x.pipeline, stage)
	return x
}

// NewPipeline to start a new pipeline.
func (x *MongoQueryBuilder) NewPipeline() *MongoQueryBuilder {
	x.pipeline = mongo.Pipeline{}

	return x
}

// Execute runs the aggregation pipeline and returns the cursor.
func (x *MongoQueryBuilder) Execute(ctx context.Context) (*mongo.Cursor, error) {
	return x.collection.Aggregate(ctx, x.pipeline)
}

// InsertOne inserts a single document into the collection.
func (x *MongoQueryBuilder) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return x.collection.InsertOne(ctx, document)
}

// InsertMany inserts multiple documents into the collection.
func (x *MongoQueryBuilder) InsertMany(ctx context.Context, documents []interface{}) (*mongo.InsertManyResult, error) {
	return x.collection.InsertMany(ctx, documents)
}

// Find runs the find operation with the given filter.
func (mq *MongoQueryBuilder) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return mq.collection.Find(ctx, filter, opts...)
}

// FindByID finds a document by its _id field.
func (x *MongoQueryBuilder) FindByID(ctx context.Context, id interface{}) *mongo.SingleResult {
	return x.collection.FindOne(ctx, bson.D{{"_id", id}})
}

// FindOne finds a single document matching the given filter.
func (x *MongoQueryBuilder) FindOne(ctx context.Context, filter bson.M) *mongo.SingleResult {
	return x.collection.FindOne(ctx, filter)
}

// FindMany finds multiple documents matching the given filter.
func (x *MongoQueryBuilder) FindMany(ctx context.Context, filter bson.D, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return x.collection.Find(ctx, filter, opts...)
}

// UpdateOne updates a single document in the collection.
func (x *MongoQueryBuilder) UpdateOne(ctx context.Context, filter bson.D, update bson.D) (*mongo.UpdateResult, error) {
	return x.collection.UpdateOne(ctx, filter, bson.D{{"$set", update}})
}

// UpdateMany updates multiple documents in the collection.
func (x *MongoQueryBuilder) UpdateMany(ctx context.Context, filter bson.D, update bson.D) (*mongo.UpdateResult, error) {
	return x.collection.UpdateMany(ctx, filter, bson.D{{"$set", update}})
}

// DeleteOne deletes a single document from the collection.
func (x *MongoQueryBuilder) DeleteOne(ctx context.Context, filter bson.M) (*mongo.DeleteResult, error) {
	return x.collection.DeleteOne(ctx, filter)
}

// DeleteMany deletes multiple documents from the collection.
func (x *MongoQueryBuilder) DeleteMany(ctx context.Context, filter bson.M) (*mongo.DeleteResult, error) {
	return x.collection.DeleteMany(ctx, filter)
}

// ReplaceOne replaces the entire document that matches the given filter.
func (x *MongoQueryBuilder) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}) (*mongo.UpdateResult, error) {
	return x.collection.ReplaceOne(ctx, filter, replacement)
}

// Count returns the number of documents that match the given filter.
func (x *MongoQueryBuilder) Count(ctx context.Context, filter bson.D) (int64, error) {
	count, err := x.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// SearchSingleField adds a search on a single field using regex.
func (x *MongoQueryBuilder) SearchSingleField(field string, searchValue string) *MongoQueryBuilder {
	x.AddStage(bson.D{{
		"$match", bson.D{
			{field, primitive.Regex{
				Pattern: searchValue,
				Options: "i",
			}},
		},
	}})

	return x
}

// SearchMultipleFields adds a search on multiple fields using regex.
func (x *MongoQueryBuilder) SearchMultipleFields(fields []string, searchValue string) *MongoQueryBuilder {
	if searchValue != "" && len(fields) > 0 {
		conditions := bson.A{}
		for _, field := range fields {
			conditions = append(conditions, bson.D{
				{field, bson.D{
					{"$regex", searchValue},
					{"$options", "i"},
				}},
			})
		}
		x.AddStage(bson.D{{
			"$match", bson.D{
				{"$or", conditions},
			},
		}})
	}
	return x
}

// Lookup adds a $lookup stage to the aggregation pipeline.
func (x *MongoQueryBuilder) Lookup(from, localField, foreignField, as string, pipeline ...bson.A) *MongoQueryBuilder {
	lookupStage := bson.D{
		{"from", from},
		{"localField", localField},
		{"foreignField", foreignField},
		{"as", as},
	}

	// Check if the pipeline is given (i.e. there is an element in the variadic argument)
	if len(pipeline) > 0 && len(pipeline[0]) > 0 {
		lookupStage = append(lookupStage, bson.E{
			Key:   "pipeline",
			Value: pipeline[0], // Take the first pipeline of the variadic
		})
	}

	x.AddStage(bson.D{{"$lookup", lookupStage}})
	return x
}

// Where adds a single condition to bson.D
func (x *MongoQueryBuilder) Where(condition *bson.D, field string, operator string, value interface{}) {
	var cond bson.E

	switch operator {
	case "=":
		cond = bson.E{Key: field, Value: value}
	case "!=":
		cond = bson.E{Key: field, Value: bson.D{{"$ne", value}}}
	case ">":
		cond = bson.E{Key: field, Value: bson.D{{"$gt", value}}}
	case ">=":
		cond = bson.E{Key: field, Value: bson.D{{"$gte", value}}}
	case "<":
		cond = bson.E{Key: field, Value: bson.D{{"$lt", value}}}
	case "<=":
		cond = bson.E{Key: field, Value: bson.D{{"$lte", value}}}
	case "in":
		cond = bson.E{Key: field, Value: bson.D{{"$in", value}}}
	case "not in":
		cond = bson.E{Key: field, Value: bson.D{{"$nin", value}}}
	case "like":
		cond = bson.E{Key: field, Value: primitive.Regex{Pattern: value.(string), Options: "i"}}
	default:
		cond = bson.E{Key: field, Value: value}
	}

	*condition = append(*condition, cond)
}

// WhereGroup acts as a callback for multiple conditions inside AddConditions
func (x *MongoQueryBuilder) WhereGroup(callback func(condition *bson.D)) bson.D {
	condition := bson.D{}
	callback(&condition)
	return condition
}

// AddConditions allows multiple WhereGroup callbacks to be added as $match
func (x *MongoQueryBuilder) AddConditions(callback func(builder *MongoQueryBuilder)) *MongoQueryBuilder {
	// Buat callback AddConditions yang memanggil WhereGroup
	callback(x)
	return x
}

// Match adds one condition to the pipeline ($match only)
func (x *MongoQueryBuilder) Match(condition bson.D) *MongoQueryBuilder {
	x.AddStage(bson.D{{
		"$match", condition,
	}})
	return x
}

// Sort adds a $sort stage to sort documents.
func (x *MongoQueryBuilder) Sort(field string, ascending bool) *MongoQueryBuilder {
	order := 1
	if !ascending {
		order = -1
	}
	x.AddStage(bson.D{{
		"$sort", bson.D{
			{field, order},
		}},
	})
	return x
}

// Paginate adds pagination support with page and per_page parameters.
func (x *MongoQueryBuilder) Paginate(page, perPage int) *MongoQueryBuilder {
	skip := (page - 1) * perPage

	x.AddStage(bson.D{{
		"$project", bson.D{
			{"data", bson.D{
				{"$slice", bson.A{
					"$data", skip, bson.D{
						{"$ifNull", bson.A{
							perPage, "$total.count",
						}},
					},
				}},
			}},
			{"page", bson.D{
				{"$literal", skip/perPage + 1},
			}},
			{"per_page", bson.D{
				{"$literal", perPage},
			}},
			{"total", bson.D{
				{"$ifNull", bson.A{"$total.count", 0}},
			}},
		},
	}})
	return x
}

// Facet adds a $facet stage to the aggregation pipeline for returning multiple outputs.
func (x *MongoQueryBuilder) Facet(facets bson.D) *MongoQueryBuilder {
	x.AddStage(bson.D{{
		"$facet", facets,
	}})
	return x
}

// Unwind adds a $unwind stage to the aggregation pipeline.
func (x *MongoQueryBuilder) Unwind(path string) *MongoQueryBuilder {
	x.AddStage(bson.D{{
		"$unwind", path,
	}})
	return x
}

// Project adds a $project stage to reshape the documents.
func (x *MongoQueryBuilder) Project(projection bson.D) *MongoQueryBuilder {
	x.AddStage(bson.D{{
		"$project", projection,
	}})
	return x
}

// CountFacet adds a counting stage in the $facet pipeline.
func (x *MongoQueryBuilder) CountFacet() *MongoQueryBuilder {
	x.Facet(bson.D{
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
	})
	return x
}

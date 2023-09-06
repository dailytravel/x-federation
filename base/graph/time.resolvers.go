package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/base/graph/model"
	"github.com/dailytravel/x/base/internal/utils"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateTime is the resolver for the createTime field.
func (r *mutationResolver) CreateTime(ctx context.Context, input model.NewTime) (*model.Time, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	item := &model.Time{
		Model: model.Model{
			Metadata:  input.Metadata,
			CreatedBy: uid,
			UpdatedBy: uid,
		},
	}

	res, err := r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	item.ID = res.InsertedID.(primitive.ObjectID)

	return item, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *mutationResolver) UpdateTime(ctx context.Context, id string, input model.UpdateTime) (*model.Time, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Fetch the existing board
	item := &model.Time{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	// Update fields based on input

	item.UpdatedBy = uid

	// Perform the update in the database
	update := bson.M{
		"$set": item,
	}
	_, err = r.db.Collection(item.Collection()).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteTime is the resolver for the deleteTime field.
func (r *mutationResolver) DeleteTime(ctx context.Context, id string) (map[string]interface{}, error) {
	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Define the filter to match the given ID
	filter := bson.M{"_id": _id}

	// Perform the delete operation
	result, err := r.db.Collection("times").DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Create and return the response
	response := map[string]interface{}{
		"status":       "success",
		"deletedCount": result.DeletedCount,
	}
	return response, nil
}

// DeleteTimes is the resolver for the deleteTimes field.
func (r *mutationResolver) DeleteTimes(ctx context.Context, ids []string) (map[string]interface{}, error) {
	// Convert the list of ID strings to a list of ObjectIDs
	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, _id)
	}

	// Define the filter to match the given IDs
	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	// Perform the delete operation
	result, err := r.db.Collection("times").DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Create and return the response
	response := map[string]interface{}{
		"status":       "success",
		"deletedCount": result.DeletedCount,
	}
	return response, nil
}

// Time is the resolver for the time field.
func (r *queryResolver) Time(ctx context.Context, id string) (*model.Time, error) {
	var item *model.Time

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}

	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found for filter %v", filter)
		}
		return nil, err
	}

	return item, nil
}

// Times is the resolver for the times field.
func (r *queryResolver) Times(ctx context.Context, args map[string]interface{}) (*model.Times, error) {
	var items []*model.Time

	opts := utils.Options(args)
	opts.SetSort(bson.M{"order": 1})
	opts.SetSort(bson.M{"created_at": -1})

	// Build the filter based on the provided arguments
	filter := bson.M{}

	// Add filters based on the arguments, if provided
	if name, ok := args["name"].(string); ok && name != "" {
		filter["name"] = name
	}

	// Create a cursor for the query
	cursor, err := r.db.Collection("times").Find(ctx, utils.Query(args), opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode documents
	for cursor.Next(ctx) {
		var item model.Time
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	// Check for cursor errors
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// You can get the total count using CountDocuments method
	count, err := r.db.Collection("times").CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &model.Times{Data: items, Count: int(count)}, nil
}

// ID is the resolver for the id field.
func (r *timeResolver) ID(ctx context.Context, obj *model.Time) (string, error) {
	return obj.ID.Hex(), nil
}

// Parent is the resolver for the parent field.
func (r *timeResolver) Parent(ctx context.Context, obj *model.Time) (*model.Time, error) {
	var item *model.Time

	if obj.Parent == nil {
		return nil, nil
	}

	filter := bson.M{"_id": obj.Parent}
	if err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found for filter %v", filter)
		}
		return nil, err
	}

	return item, nil
}

// Start is the resolver for the start field.
func (r *timeResolver) Start(ctx context.Context, obj *model.Time) (string, error) {
	panic(fmt.Errorf("not implemented: Start - start"))
}

// End is the resolver for the end field.
func (r *timeResolver) End(ctx context.Context, obj *model.Time) (string, error) {
	panic(fmt.Errorf("not implemented: End - end"))
}

// Metadata is the resolver for the metadata field.
func (r *timeResolver) Metadata(ctx context.Context, obj *model.Time) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *timeResolver) CreatedAt(ctx context.Context, obj *model.Time) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *timeResolver) UpdatedAt(ctx context.Context, obj *model.Time) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *timeResolver) CreatedBy(ctx context.Context, obj *model.Time) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	return pointer.String(obj.CreatedBy.Hex()), nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *timeResolver) UpdatedBy(ctx context.Context, obj *model.Time) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	return pointer.String(obj.UpdatedBy.Hex()), nil
}

// Time returns TimeResolver implementation.
func (r *Resolver) Time() TimeResolver { return &timeResolver{r} }

type timeResolver struct{ *Resolver }

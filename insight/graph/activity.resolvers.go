package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/insight/graph/model"
	"github.com/dailytravel/x/insight/internal/utils"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ID is the resolver for the id field.
func (r *activityResolver) ID(ctx context.Context, obj *model.Activity) (string, error) {
	return obj.ID.Hex(), nil
}

// Metadata is the resolver for the metadata field.
func (r *activityResolver) Metadata(ctx context.Context, obj *model.Activity) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// UID is the resolver for the uid field.
func (r *activityResolver) UID(ctx context.Context, obj *model.Activity) (string, error) {
	return obj.ID.Hex(), nil
}

// Object is the resolver for the object field.
func (r *activityResolver) Object(ctx context.Context, obj *model.Activity) (string, error) {
	panic(fmt.Errorf("not implemented: Object - object"))
}

// Timestamp is the resolver for the timestamp field.
func (r *activityResolver) Timestamp(ctx context.Context, obj *model.Activity) (string, error) {
	return time.Unix(int64(obj.Timestamp.T), 0).Format(time.RFC3339), nil
}

// DeleteActivity is the resolver for the deleteActivity field.
func (r *mutationResolver) DeleteActivity(ctx context.Context, id string) (*bool, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res, err := r.db.Collection("activities").DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, fmt.Errorf("error deleting activity: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("activity not found")
	}

	return pointer.True(), nil
}

// DeleteActivities is the resolver for the deleteActivities field.
func (r *mutationResolver) DeleteActivities(ctx context.Context, filter map[string]interface{}) (*bool, error) {
	res, err := r.db.Collection("activities").DeleteMany(ctx, utils.Filter(filter))
	if err != nil {
		return nil, fmt.Errorf("error deleting activity: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("activities not found")
	}

	return pointer.True(), nil
}

// Activity is the resolver for the activity field.
func (r *queryResolver) Activity(ctx context.Context, id string) (*model.Activity, error) {
	panic(fmt.Errorf("not implemented: Activities - activities"))
}

// Activities is the resolver for the activities field.
func (r *queryResolver) Activities(ctx context.Context, args map[string]interface{}) (*model.Activities, error) {
	panic(fmt.Errorf("not implemented: Activities - activities"))
}

// Activity returns ActivityResolver implementation.
func (r *Resolver) Activity() ActivityResolver { return &activityResolver{r} }

type activityResolver struct{ *Resolver }

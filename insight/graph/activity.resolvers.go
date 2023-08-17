package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/insight/graph/model"
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

// Target is the resolver for the target field.
func (r *activityResolver) Target(ctx context.Context, obj *model.Activity) (string, error) {
	return obj.Target.Hex(), nil
}

// CreatedAt is the resolver for the created_at field.
func (r *activityResolver) CreatedAt(ctx context.Context, obj *model.Activity) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
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

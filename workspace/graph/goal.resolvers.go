package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/workspace/graph/model"
	"github.com/dailytravel/x/workspace/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ID is the resolver for the id field.
func (r *goalResolver) ID(ctx context.Context, obj *model.Goal) (string, error) {
	return obj.ID.Hex(), nil
}

// Metadata is the resolver for the metadata field.
func (r *goalResolver) Metadata(ctx context.Context, obj *model.Goal) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *goalResolver) CreatedAt(ctx context.Context, obj *model.Goal) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *goalResolver) UpdatedAt(ctx context.Context, obj *model.Goal) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// Parent is the resolver for the parent field.
func (r *goalResolver) Parent(ctx context.Context, obj *model.Goal) (*model.Goal, error) {
	var parent *model.Goal

	if obj.Parent != nil {
		return nil, nil
	}

	if err := r.db.Collection("goals").FindOne(ctx, bson.M{"_id": obj.Parent}).Decode(&parent); err != nil {
		return nil, err
	}

	return parent, nil
}

// Time is the resolver for the time field.
func (r *goalResolver) Time(ctx context.Context, obj *model.Goal) (*model.Time, error) {
	var time *model.Time

	if err := r.db.Collection("times").FindOne(ctx, bson.M{"_id": obj.Time}).Decode(&time); err != nil {
		return nil, err
	}

	return time, nil
}

// UID is the resolver for the uid field.
func (r *goalResolver) UID(ctx context.Context, obj *model.Goal) (string, error) {
	return obj.ID.Hex(), nil
}

// Organization is the resolver for the organization field.
func (r *goalResolver) Organization(ctx context.Context, obj *model.Goal) (*string, error) {
	if obj.Organization == nil {
		return nil, nil
	}

	organization := obj.Organization.Hex()

	return &organization, nil
}

// CreatedBy is the resolver for the created_by field.
func (r *goalResolver) CreatedBy(ctx context.Context, obj *model.Goal) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	createdBy := obj.CreatedBy.Hex()

	return &createdBy, nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *goalResolver) UpdatedBy(ctx context.Context, obj *model.Goal) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	updatedBy := obj.UpdatedBy.Hex()

	return &updatedBy, nil
}

// CreateGoal is the resolver for the createGoal field.
func (r *mutationResolver) CreateGoal(ctx context.Context, input model.NewGoal) (*model.Goal, error) {
	panic(fmt.Errorf("not implemented: CreateGoal - createGoal"))
}

// UpdateGoal is the resolver for the updateGoal field.
func (r *mutationResolver) UpdateGoal(ctx context.Context, id string, input model.UpdateGoal) (*model.Goal, error) {
	panic(fmt.Errorf("not implemented: UpdateGoal - updateGoal"))
}

// DeleteGoal is the resolver for the deleteGoal field.
func (r *mutationResolver) DeleteGoal(ctx context.Context, id string) (map[string]interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}

	result, err := r.db.Collection("goals").DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return map[string]interface{}{
			"deleted": false,
			"error":   "goal not found",
		}, nil
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}

// DeleteGoals is the resolver for the deleteGoals field.
func (r *mutationResolver) DeleteGoals(ctx context.Context, ids []string) (map[string]interface{}, error) {
	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, _id)
	}

	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	result, err := r.db.Collection("goals").DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return map[string]interface{}{
			"deleted": false,
			"error":   "no goals were deleted",
		}, nil
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}

// Goal is the resolver for the goal field.
func (r *queryResolver) Goal(ctx context.Context, id string) (*model.Goal, error) {
	var item *model.Goal

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Collection("goals").FindOne(ctx, bson.M{"_id": _id}).Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// Goals is the resolver for the goals field.
func (r *queryResolver) Goals(ctx context.Context, args map[string]interface{}) (*model.Goals, error) {
	var items []*model.Goal

	cursor, err := r.db.Collection("goals").Find(ctx, utils.Query(args), utils.Options(args))
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var item *model.Goal

		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	count, err := r.db.Collection("goals").CountDocuments(ctx, utils.Query(args))
	if err != nil {
		return nil, err
	}

	return &model.Goals{Data: items, Count: int(count)}, nil
}

// Goal returns GoalResolver implementation.
func (r *Resolver) Goal() GoalResolver { return &goalResolver{r} }

type goalResolver struct{ *Resolver }

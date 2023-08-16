package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/configuration/graph/model"
	"github.com/dailytravel/x/configuration/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateTimezone is the resolver for the createTimezone field.
func (r *mutationResolver) CreateTimezone(ctx context.Context, input model.NewTimezone) (*model.Timezone, error) {
	panic(fmt.Errorf("not implemented: CreateTimezone - createTimezone"))
}

// UpdateTimezone is the resolver for the updateTimezone field.
func (r *mutationResolver) UpdateTimezone(ctx context.Context, input model.UpdateTimezone) (*model.Timezone, error) {
	panic(fmt.Errorf("not implemented: UpdateTimezone - updateTimezone"))
}

// ImportTimezones is the resolver for the importTimezones field.
func (r *mutationResolver) ImportTimezones(ctx context.Context, url string) ([]*model.Timezone, error) {
	panic(fmt.Errorf("not implemented: ImportTimezones - importTimezones"))
}

// DeleteTimezone is the resolver for the deleteTimezone field.
func (r *mutationResolver) DeleteTimezone(ctx context.Context, id string) (map[string]interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res, err := r.db.Collection("timezones").DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, fmt.Errorf("error deleting log: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("log not found")
	}

	return map[string]interface{}{
		"status": "success",
	}, nil
}

// DeleteTimezones is the resolver for the deleteTimezones field.
func (r *mutationResolver) DeleteTimezones(ctx context.Context, ids []string) (map[string]interface{}, error) {
	var _ids []primitive.ObjectID

	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		_ids = append(_ids, _id)
	}

	res, err := r.db.Collection("timezones").DeleteMany(ctx, bson.M{"_id": bson.M{"$in": _ids}})
	if err != nil {
		return nil, fmt.Errorf("error deleting log: %v", err)
	}

	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("log not found")
	}

	return map[string]interface{}{
		"status": "success",
	}, nil
}

// Timezones is the resolver for the timezones field.
func (r *queryResolver) Timezones(ctx context.Context, args map[string]interface{}) (*model.Timezones, error) {
	var items []*model.Timezone
	//find all items
	cur, err := r.db.Collection("timezones").Find(ctx, utils.Query(args), utils.Options(args))
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Timezone
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("timezones").CountDocuments(ctx, utils.Query(args), nil)
	if err != nil {
		return nil, err
	}

	return &model.Timezones{
		Count: int(count),
		Data:  items,
	}, nil
}

// Timezone is the resolver for the timezone field.
func (r *queryResolver) Timezone(ctx context.Context, id string) (*model.Timezone, error) {
	var item *model.Timezone
	col := r.db.Collection(item.Collection())

	filter := bson.M{"_id": id}

	err := col.FindOne(ctx, filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found for filter %v", filter)
		}
		return nil, err
	}

	return item, nil
}

// ID is the resolver for the id field.
func (r *timezoneResolver) ID(ctx context.Context, obj *model.Timezone) (string, error) {
	return obj.ID.Hex(), nil
}

// Metadata is the resolver for the metadata field.
func (r *timezoneResolver) Metadata(ctx context.Context, obj *model.Timezone) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *timezoneResolver) CreatedAt(ctx context.Context, obj *model.Timezone) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *timezoneResolver) UpdatedAt(ctx context.Context, obj *model.Timezone) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *timezoneResolver) CreatedBy(ctx context.Context, obj *model.Timezone) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	createdBy := obj.CreatedBy.Hex()

	return &createdBy, nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *timezoneResolver) UpdatedBy(ctx context.Context, obj *model.Timezone) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	updatedBy := obj.UpdatedBy.Hex()

	return &updatedBy, nil
}

// Timezone returns TimezoneResolver implementation.
func (r *Resolver) Timezone() TimezoneResolver { return &timezoneResolver{r} }

type timezoneResolver struct{ *Resolver }
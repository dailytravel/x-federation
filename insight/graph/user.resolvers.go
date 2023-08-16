package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/dailytravel/x/insight/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Logs is the resolver for the logs field.
func (r *userResolver) Logs(ctx context.Context, obj *model.User) ([]*model.Log, error) {
	var items []*model.Log

	uid, err := primitive.ObjectIDFromHex(obj.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"uid": uid}
	//find all items
	cur, err := r.db.Collection("logs").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Log
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// Activities is the resolver for the activities field.
func (r *userResolver) Activities(ctx context.Context, obj *model.User) ([]*model.Activity, error) {
	var items []*model.Activity

	uid, err := primitive.ObjectIDFromHex(obj.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"uid": uid}
	//find all items
	cur, err := r.db.Collection("activities").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Activity
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
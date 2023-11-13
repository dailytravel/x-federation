package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/account/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User is the resolver for the user field.
func (r *collaboratorResolver) User(ctx context.Context, obj *model.Collaborator) (*model.User, error) {
	_id, err := primitive.ObjectIDFromHex(obj.UID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert UID to ObjectID: %w", err)
	}

	var item *model.User

	filter := bson.M{"_id": _id}
	if err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Return nil if no user is found, rather than an error.
		}
		return nil, fmt.Errorf("failed to fetch user from database: %w", err)
	}

	return item, nil
}

// Collaborator returns CollaboratorResolver implementation.
func (r *Resolver) Collaborator() CollaboratorResolver { return &collaboratorResolver{r} }

type collaboratorResolver struct{ *Resolver }

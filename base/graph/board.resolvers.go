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
	"github.com/dailytravel/x/base/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ID is the resolver for the id field.
func (r *boardResolver) ID(ctx context.Context, obj *model.Board) (string, error) {
	return obj.ID.Hex(), nil
}

// Portfolio is the resolver for the portfolio field.
func (r *boardResolver) Portfolio(ctx context.Context, obj *model.Board) (*model.Portfolio, error) {
	var item *model.Portfolio

	if err := r.db.Collection("portfolios").FindOne(ctx, bson.M{"_id": obj.Portfolio}).Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// DueDate is the resolver for the due_date field.
func (r *boardResolver) DueDate(ctx context.Context, obj *model.Board) (*string, error) {
	if obj.DueDate == nil {
		return nil, nil
	}

	dueDate := obj.DueDate.Time().Format(time.RFC3339)
	return &dueDate, nil
}

// Metadata is the resolver for the metadata field.
func (r *boardResolver) Metadata(ctx context.Context, obj *model.Board) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *boardResolver) CreatedAt(ctx context.Context, obj *model.Board) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *boardResolver) UpdatedAt(ctx context.Context, obj *model.Board) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// Lists is the resolver for the lists field.
func (r *boardResolver) Lists(ctx context.Context, obj *model.Board) ([]*model.List, error) {
	var items []*model.List

	opts := options.Find().SetSort(bson.M{"order": 1})

	cursor, err := r.db.Collection("lists").Find(ctx, bson.M{"board": obj.ID}, opts)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(context.Background(), &items); err != nil {
		return nil, err
	}

	return items, nil
}

// UID is the resolver for the uid field.
func (r *boardResolver) UID(ctx context.Context, obj *model.Board) (string, error) {
	return obj.UID.Hex(), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *boardResolver) CreatedBy(ctx context.Context, obj *model.Board) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	createdBy := obj.CreatedBy.Hex()

	return &createdBy, nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *boardResolver) UpdatedBy(ctx context.Context, obj *model.Board) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	updatedBy := obj.UpdatedBy.Hex()

	return &updatedBy, nil
}

// CreateBoard is the resolver for the createBoard field.
func (r *mutationResolver) CreateBoard(ctx context.Context, input model.NewBoard) (*model.Board, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	item := &model.Board{
		UID:         *uid,
		Title:       input.Title,
		Type:        input.Type,
		Description: input.Description,
		IsTemplate:  input.IsTemplate,
		Order:       input.Order,
		Starred:     input.Starred,
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

// UpdateBoard is the resolver for the updateBoard field.
func (r *mutationResolver) UpdateBoard(ctx context.Context, id string, input model.UpdateBoard) (*model.Board, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Fetch the existing board
	item := &model.Board{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	// Update fields based on input
	if input.Title != nil {
		item.Title = *input.Title
	}
	if input.Description != nil {
		item.Description = input.Description
	}
	if input.IsTemplate != nil {
		item.IsTemplate = *input.IsTemplate
	}
	if input.Order != nil {
		item.Order = *input.Order
	}
	if input.Starred != nil {
		item.Starred = *input.Starred
	}
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

// DeleteBoard is the resolver for the deleteBoard field.
func (r *mutationResolver) DeleteBoard(ctx context.Context, id string) (map[string]interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}

	result, err := r.db.Collection("boards").DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return map[string]interface{}{
			"deleted": false,
			"error":   "board not found",
		}, nil
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}

// DeleteBoards is the resolver for the deleteBoards field.
func (r *mutationResolver) DeleteBoards(ctx context.Context, ids []string) (map[string]interface{}, error) {
	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, _id)
	}

	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	result, err := r.db.Collection("boards").DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return map[string]interface{}{
			"deleted": false,
			"error":   "no boards were deleted",
		}, nil
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}

// Board is the resolver for the board field.
func (r *queryResolver) Board(ctx context.Context, id string) (*model.Board, error) {
	var item *model.Board

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}

	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found for filter %v", filter)
		}
		return nil, err
	}

	return item, nil
}

// Boards is the resolver for the boards field.
func (r *queryResolver) Boards(ctx context.Context, args map[string]interface{}) (*model.Boards, error) {
	var items []*model.Board

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
	cursor, err := r.db.Collection("boards").Find(ctx, utils.Query(args), opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode documents
	for cursor.Next(ctx) {
		var board model.Board
		if err := cursor.Decode(&board); err != nil {
			return nil, err
		}
		items = append(items, &board)
	}

	// Check for cursor errors
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// You can get the total count using CountDocuments method
	count, err := r.db.Collection("boards").CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &model.Boards{Data: items, Count: int(count)}, nil
}

// Board returns BoardResolver implementation.
func (r *Resolver) Board() BoardResolver { return &boardResolver{r} }

type boardResolver struct{ *Resolver }
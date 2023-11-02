package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"log"
	"time"

	"github.com/dailytravel/x/base/graph/model"
	"github.com/dailytravel/x/base/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ID is the resolver for the id field.
func (r *listResolver) ID(ctx context.Context, obj *model.List) (string, error) {
	return obj.ID.Hex(), nil
}

// Board is the resolver for the board field.
func (r *listResolver) Board(ctx context.Context, obj *model.List) (*model.Board, error) {
	var item *model.Board

	if err := r.db.Collection("boards").FindOne(ctx, bson.M{"_id": obj.Board}).Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// Metadata is the resolver for the metadata field.
func (r *listResolver) Metadata(ctx context.Context, obj *model.List) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Tasks is the resolver for the tasks field.
func (r *listResolver) Tasks(ctx context.Context, obj *model.List) ([]*model.Task, error) {
	if len(obj.Tasks) == 0 {
		return []*model.Task{}, nil
	}

	var items []*model.Task

	filter := bson.M{"_id": bson.M{"$in": obj.Tasks}, "status": bson.M{"$ne": "deleted"}}
	cursor, err := r.db.Collection("tasks").Find(ctx, filter, nil)
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
func (r *listResolver) UID(ctx context.Context, obj *model.List) (string, error) {
	return obj.UID.Hex(), nil
}

// Created is the resolver for the created field.
func (r *listResolver) Created(ctx context.Context, obj *model.List) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *listResolver) Updated(ctx context.Context, obj *model.List) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// CreateList is the resolver for the createList field.
func (r *mutationResolver) CreateList(ctx context.Context, input model.NewList) (*model.List, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		log.Printf("error getting uid: %v", err)
		return nil, err
	}

	log.Println("CreateList")

	board, err := primitive.ObjectIDFromHex(input.Board)
	if err != nil {
		return nil, err
	}

	item := &model.List{
		UID:   *uid,
		Board: board,
		Name:  input.Name,
	}

	if input.Order != nil {
		item.Order = *input.Order
	}

	if input.Tasks != nil {
		var taskIDs []primitive.ObjectID // Create a slice to collect converted task IDs

		for _, task := range input.Tasks {
			_id, err := primitive.ObjectIDFromHex(*task)
			if err != nil {
				return nil, err
			}
			taskIDs = append(taskIDs, _id) // Collect the converted IDs
		}

		item.Tasks = taskIDs // Assign the collected IDs to item.Tasks
	}

	if input.Metadata != nil {
		if item.Metadata == nil {
			item.Metadata = make(map[string]interface{})
		}

		for k, v := range input.Metadata {
			item.Metadata[k] = v
		}
	}

	res, err := r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	item.ID = res.InsertedID.(primitive.ObjectID)

	return item, nil
}

// UpdateList is the resolver for the updateList field.
func (r *mutationResolver) UpdateList(ctx context.Context, id string, input model.UpdateList) (*model.List, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Fetch the existing list
	existingList := &model.List{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(existingList.Collection()).FindOne(ctx, filter).Decode(existingList)
	if err != nil {
		return nil, err
	}

	// Update fields based on input
	if input.Name != nil {
		existingList.Name = *input.Name
	}
	if input.Order != nil {
		existingList.Order = *input.Order
	}

	// Perform the update in the database
	update := bson.M{
		"$set": existingList,
	}
	_, err = r.db.Collection(existingList.Collection()).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return existingList, nil
}

// DeleteList is the resolver for the deleteList field.
func (r *mutationResolver) DeleteList(ctx context.Context, id string) (map[string]interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}

	result, err := r.db.Collection("lists").DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return map[string]interface{}{
			"deleted": false,
			"error":   "list not found",
		}, nil
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}

// DeleteLists is the resolver for the deleteLists field.
func (r *mutationResolver) DeleteLists(ctx context.Context, ids []string) (map[string]interface{}, error) {
	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, _id)
	}

	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	result, err := r.db.Collection("lists").DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return map[string]interface{}{
			"deleted": false,
			"error":   "no lists were deleted",
		}, nil
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}

// List is the resolver for the list field.
func (r *queryResolver) List(ctx context.Context, id string) (*model.List, error) {
	var item *model.List
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Collection("lists").FindOne(ctx, bson.M{"_id": _id}).Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// Lists is the resolver for the lists field.
func (r *queryResolver) Lists(ctx context.Context, board string) (*model.Lists, error) {
	var items []*model.List
	_id, err := primitive.ObjectIDFromHex(board)
	if err != nil {
		return nil, err
	}

	opts := options.Find().SetSort(bson.M{"order": 1})
	opts.SetSort(bson.M{"created": -1})
	filter := bson.M{"board": _id}

	//find all items
	cur, err := r.db.Collection("lists").Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {

		var item *model.List
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("lists").CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &model.Lists{
		Count: int(count),
		Data:  items,
	}, nil
}

// List returns ListResolver implementation.
func (r *Resolver) List() ListResolver { return &listResolver{r} }

type listResolver struct{ *Resolver }

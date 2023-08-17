package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/sales/auth"
	"github.com/dailytravel/x/sales/graph/model"
	"github.com/dailytravel/x/sales/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateTransaction is the resolver for the createTransaction field.
func (r *mutationResolver) CreateTransaction(ctx context.Context, input *model.NewTransaction) (*model.Transaction, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	item := &model.Transaction{
		Model: model.Model{
			CreatedBy: uid,
			UpdatedBy: uid,
			Metadata:  input.Metadata,
		},
	}

	// Set the fields from the input
	_, err = r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateTransaction is the resolver for the updateTransaction field.
func (r *mutationResolver) UpdateTransaction(ctx context.Context, id string, input *model.UpdateTransaction) (*model.Transaction, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Create an update document with the fields to be updated
	item := &model.Transaction{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	if input.Metadata != nil {
		for k, v := range input.Metadata {
			item.Metadata[k] = v
		}
	}

	// Update the updated_by and updated_at fields
	item.UpdatedBy = uid

	// Perform the update in the database
	res, err := r.db.Collection(item.Collection()).UpdateOne(ctx, filter, item)
	if err != nil {
		return nil, err
	}

	// Check if the coupon was actually updated
	if res.ModifiedCount == 0 {
		return nil, fmt.Errorf("no coupon was updated")
	}

	return item, nil
}

// DeleteTransaction is the resolver for the deleteTransaction field.
func (r *mutationResolver) DeleteTransaction(ctx context.Context, id string) (map[string]interface{}, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	// Convert the ID string to an ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Define the filter to match the given ID
	filter := bson.M{"_id": _id}

	// Define the update to mark the record as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"deleted_by": uid,
			"status":     "deleted",
			"updated_by": uid,
			"updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Perform the update operation in the database
	result, err := r.db.Collection("transactions").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("transaction not found")
	}

	return map[string]interface{}{"status": "success", "deletedCount": result.ModifiedCount}, nil
}

// DeleteTransactions is the resolver for the deleteTransactions field.
func (r *mutationResolver) DeleteTransactions(ctx context.Context, ids []string) (map[string]interface{}, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	// Convert the list of ID strings to ObjectIDs
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

	// Define the update to mark records as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"deleted_by": uid,
			"status":     "deleted",
			"updated_by": uid,
			"updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Perform the update operation in the database
	result, err := r.db.Collection("transactions").UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "deletedCount": result.ModifiedCount}, nil
}

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context, args map[string]interface{}) (*model.Transactions, error) {
	var items []*model.Transaction
	//find all items
	cur, err := r.db.Collection("transactions").Find(ctx, utils.Query(args), utils.Options(args))
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Transaction
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("transactions").CountDocuments(ctx, utils.Query(args), nil)
	if err != nil {
		return nil, err
	}

	return &model.Transactions{
		Count: int(count),
		Data:  items,
	}, nil
}

// Transaction is the resolver for the transaction field.
func (r *queryResolver) Transaction(ctx context.Context, id string) (*model.Transaction, error) {
	var item *model.Transaction

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}
	if err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("document not found")
		}
		return nil, err
	}

	return item, nil
}

// ID is the resolver for the id field.
func (r *transactionResolver) ID(ctx context.Context, obj *model.Transaction) (string, error) {
	return obj.ID.Hex(), nil
}

// Metadata is the resolver for the metadata field.
func (r *transactionResolver) Metadata(ctx context.Context, obj *model.Transaction) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Date is the resolver for the date field.
func (r *transactionResolver) Date(ctx context.Context, obj *model.Transaction) (string, error) {
	return obj.Date.Time().String(), nil
}

// CreatedAt is the resolver for the created_at field.
func (r *transactionResolver) CreatedAt(ctx context.Context, obj *model.Transaction) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *transactionResolver) UpdatedAt(ctx context.Context, obj *model.Transaction) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// UID is the resolver for the uid field.
func (r *transactionResolver) UID(ctx context.Context, obj *model.Transaction) (string, error) {
	return obj.UID.Hex(), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *transactionResolver) CreatedBy(ctx context.Context, obj *model.Transaction) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	createdBy := obj.CreatedBy.Hex()

	return &createdBy, nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *transactionResolver) UpdatedBy(ctx context.Context, obj *model.Transaction) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	updatedBy := obj.UpdatedBy.Hex()

	return &updatedBy, nil
}

// Transaction returns TransactionResolver implementation.
func (r *Resolver) Transaction() TransactionResolver { return &transactionResolver{r} }

type transactionResolver struct{ *Resolver }

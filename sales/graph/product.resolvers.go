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

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	item := &model.Product{
		Locale:      input.Locale,
		Name:        bson.M{input.Locale: input.Name},
		Description: bson.M{input.Locale: input.Description},
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

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, input model.UpdateProduct) (*model.Product, error) {
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
	item := &model.Product{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	if input.Type != nil {
		item.Type = *input.Type
	}

	if input.Name != nil {
		item.Name[item.Locale] = *input.Name
	}

	if input.Description != nil {
		item.Description[item.Locale] = *input.Description
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

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (map[string]interface{}, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
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
	result, err := r.db.Collection("products").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// Return a response indicating the deleted product count
	response := map[string]interface{}{"status": "success", "deletedCount": result.ModifiedCount}
	return response, nil
}

// DeleteProducts is the resolver for the deleteProducts field.
func (r *mutationResolver) DeleteProducts(ctx context.Context, ids []string) (map[string]interface{}, error) {
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
	result, err := r.db.Collection("products").UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// Return a response indicating the deleted product count
	response := map[string]interface{}{"status": "success", "deletedCount": result.ModifiedCount}
	return response, nil
}

// ID is the resolver for the id field.
func (r *productResolver) ID(ctx context.Context, obj *model.Product) (string, error) {
	return obj.ID.Hex(), nil
}

// Name is the resolver for the name field.
func (r *productResolver) Name(ctx context.Context, obj *model.Product) (string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)

	// Try to retrieve the name for the requested locale
	if name, ok := obj.Name[*locale].(string); ok {
		return name, nil
	}

	// If the name is not found for the requested locale,
	// fallback to the taxonomy's default locale
	if name, ok := obj.Name[obj.Locale].(string); ok {
		return name, nil
	}

	// Return an error if the name is not found for any locale
	return "", errors.New("Name not found for any locale")
}

// Description is the resolver for the description field.
func (r *productResolver) Description(ctx context.Context, obj *model.Product) (*string, error) {
	// Get the locale from the context
	locale := auth.Locale(ctx)

	// Try to retrieve the description for the requested locale
	if description, ok := obj.Description[*locale].(string); ok {
		return &description, nil
	}

	// If the description is not found for the requested locale,
	// fallback to the taxonomy's default locale
	if description, ok := obj.Description[obj.Locale].(string); ok {
		return &description, nil
	}

	// Return an error if the name is not found for any locale
	return nil, errors.New("Description not found for any locale")
}

// Prices is the resolver for the prices field.
func (r *productResolver) Prices(ctx context.Context, obj *model.Product) ([]*model.Price, error) {
	panic(fmt.Errorf("not implemented: Prices - prices"))
}

// Inventory is the resolver for the inventory field.
func (r *productResolver) Inventory(ctx context.Context, obj *model.Product) ([]*model.Inventory, error) {
	panic(fmt.Errorf("not implemented: Inventory - inventory"))
}

// Metadata is the resolver for the metadata field.
func (r *productResolver) Metadata(ctx context.Context, obj *model.Product) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *productResolver) CreatedAt(ctx context.Context, obj *model.Product) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *productResolver) UpdatedAt(ctx context.Context, obj *model.Product) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// UID is the resolver for the uid field.
func (r *productResolver) UID(ctx context.Context, obj *model.Product) (string, error) {
	return obj.ID.Hex(), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *productResolver) CreatedBy(ctx context.Context, obj *model.Product) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	createdBy := obj.CreatedBy.Hex()

	return &createdBy, nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *productResolver) UpdatedBy(ctx context.Context, obj *model.Product) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	updatedBy := obj.UpdatedBy.Hex()

	return &updatedBy, nil
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	var item *model.Product

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

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, args map[string]interface{}) (*model.Products, error) {
	panic(fmt.Errorf("not implemented: Products - products"))
}

// Product returns ProductResolver implementation.
func (r *Resolver) Product() ProductResolver { return &productResolver{r} }

type productResolver struct{ *Resolver }

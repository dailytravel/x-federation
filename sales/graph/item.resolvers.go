package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/sales/graph/model"
)

// ID is the resolver for the id field.
func (r *itemResolver) ID(ctx context.Context, obj *model.Item) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// UID is the resolver for the uid field.
func (r *itemResolver) UID(ctx context.Context, obj *model.Item) (string, error) {
	panic(fmt.Errorf("not implemented: UID - uid"))
}

// Package is the resolver for the package field.
func (r *itemResolver) Package(ctx context.Context, obj *model.Item) (*model.Package, error) {
	panic(fmt.Errorf("not implemented: Package - package"))
}

// Name is the resolver for the name field.
func (r *itemResolver) Name(ctx context.Context, obj *model.Item) (string, error) {
	panic(fmt.Errorf("not implemented: Name - name"))
}

// Description is the resolver for the description field.
func (r *itemResolver) Description(ctx context.Context, obj *model.Item) (*string, error) {
	panic(fmt.Errorf("not implemented: Description - description"))
}

// Quantity is the resolver for the quantity field.
func (r *itemResolver) Quantity(ctx context.Context, obj *model.Item) (int, error) {
	panic(fmt.Errorf("not implemented: Quantity - quantity"))
}

// Metadata is the resolver for the metadata field.
func (r *itemResolver) Metadata(ctx context.Context, obj *model.Item) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: Metadata - metadata"))
}

// Status is the resolver for the status field.
func (r *itemResolver) Status(ctx context.Context, obj *model.Item) (*string, error) {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// Created is the resolver for the created field.
func (r *itemResolver) Created(ctx context.Context, obj *model.Item) (string, error) {
	panic(fmt.Errorf("not implemented: Created - created"))
}

// Updated is the resolver for the updated field.
func (r *itemResolver) Updated(ctx context.Context, obj *model.Item) (string, error) {
	panic(fmt.Errorf("not implemented: Updated - updated"))
}

// CreateItem is the resolver for the createItem field.
func (r *mutationResolver) CreateItem(ctx context.Context, input model.NewItem) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: CreateItem - createItem"))
}

// UpdateItem is the resolver for the updateItem field.
func (r *mutationResolver) UpdateItem(ctx context.Context, id string, input model.UpdateItem) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: UpdateItem - updateItem"))
}

// DeleteItem is the resolver for the deleteItem field.
func (r *mutationResolver) DeleteItem(ctx context.Context, id string) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: DeleteItem - deleteItem"))
}

// DeleteItems is the resolver for the deleteItems field.
func (r *mutationResolver) DeleteItems(ctx context.Context, filter map[string]interface{}) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: DeleteItems - deleteItems"))
}

// Items is the resolver for the items field.
func (r *queryResolver) Items(ctx context.Context, filter map[string]interface{}, project map[string]interface{}, sort map[string]interface{}, collation map[string]interface{}, limit *int, skip *int) (*model.Items, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Item is the resolver for the item field.
func (r *queryResolver) Item(ctx context.Context, id string) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: Item - item"))
}

// Item returns ItemResolver implementation.
func (r *Resolver) Item() ItemResolver { return &itemResolver{r} }

type itemResolver struct{ *Resolver }

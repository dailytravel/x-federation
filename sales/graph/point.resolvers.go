package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/sales/graph/model"
)

// CreatePoint is the resolver for the createPoint field.
func (r *mutationResolver) CreatePoint(ctx context.Context, input model.NewPoint) (*model.Point, error) {
	panic(fmt.Errorf("not implemented: CreatePoint - createPoint"))
}

// UpdatePoint is the resolver for the updatePoint field.
func (r *mutationResolver) UpdatePoint(ctx context.Context, input model.UpdatePoint) (*model.Point, error) {
	panic(fmt.Errorf("not implemented: UpdatePoint - updatePoint"))
}

// DeletePoint is the resolver for the deletePoint field.
func (r *mutationResolver) DeletePoint(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeletePoint - deletePoint"))
}

// DeletePoints is the resolver for the deletePoints field.
func (r *mutationResolver) DeletePoints(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeletePoints - deletePoints"))
}

// Points is the resolver for the points field.
func (r *queryResolver) Points(ctx context.Context, args map[string]interface{}) (*model.Points, error) {
	panic(fmt.Errorf("not implemented: Points - points"))
}

// Point is the resolver for the point field.
func (r *queryResolver) Point(ctx context.Context, id string) (*model.Point, error) {
	panic(fmt.Errorf("not implemented: Point - point"))
}
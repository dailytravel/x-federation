package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/account/graph/model"
)

// CreatePermission is the resolver for the createPermission field.
func (r *mutationResolver) CreatePermission(ctx context.Context, input model.NewPermission) (*model.Permission, error) {
	panic(fmt.Errorf("not implemented: CreatePermission - createPermission"))
}

// UpdatePermission is the resolver for the updatePermission field.
func (r *mutationResolver) UpdatePermission(ctx context.Context, id string, input model.UpdatePermission) (*model.Permission, error) {
	panic(fmt.Errorf("not implemented: UpdatePermission - updatePermission"))
}

// DeletePermission is the resolver for the deletePermission field.
func (r *mutationResolver) DeletePermission(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeletePermission - deletePermission"))
}

// DeletePermissions is the resolver for the deletePermissions field.
func (r *mutationResolver) DeletePermissions(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeletePermissions - deletePermissions"))
}

// Permissions is the resolver for the permissions field.
func (r *queryResolver) Permissions(ctx context.Context, args map[string]interface{}) (*model.Permissions, error) {
	panic(fmt.Errorf("not implemented: Permissions - permissions"))
}

// Permission is the resolver for the permission field.
func (r *queryResolver) Permission(ctx context.Context, id string) (*model.Permission, error) {
	panic(fmt.Errorf("not implemented: Permission - permission"))
}
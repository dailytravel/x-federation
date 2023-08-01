package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/account/graph/model"
)

// CreateInvitation is the resolver for the createInvitation field.
func (r *mutationResolver) CreateInvitation(ctx context.Context, input model.NewInvitation) (*model.Invitation, error) {
	panic(fmt.Errorf("not implemented: CreateInvitation - createInvitation"))
}

// UpdateInvitation is the resolver for the updateInvitation field.
func (r *mutationResolver) UpdateInvitation(ctx context.Context, id string, input model.UpdateInvitation) (*model.Invitation, error) {
	panic(fmt.Errorf("not implemented: UpdateInvitation - updateInvitation"))
}

// DeleteInvitation is the resolver for the deleteInvitation field.
func (r *mutationResolver) DeleteInvitation(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteInvitation - deleteInvitation"))
}

// DeleteInvitations is the resolver for the deleteInvitations field.
func (r *mutationResolver) DeleteInvitations(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteInvitations - deleteInvitations"))
}

// Invitations is the resolver for the invitations field.
func (r *queryResolver) Invitations(ctx context.Context, args map[string]interface{}) (*model.Invitations, error) {
	panic(fmt.Errorf("not implemented: Invitations - invitations"))
}

// Invitation is the resolver for the invitation field.
func (r *queryResolver) Invitation(ctx context.Context, id string) (*model.Invitation, error) {
	panic(fmt.Errorf("not implemented: Invitation - invitation"))
}

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/dailytravel/x/notification/graph/model"
)

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}

// UpdateComment is the resolver for the updateComment field.
func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input model.UpdateComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: UpdateComment - updateComment"))
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteComment - deleteComment"))
}

// DeleteComments is the resolver for the deleteComments field.
func (r *mutationResolver) DeleteComments(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteComments - deleteComments"))
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, args map[string]interface{}) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context, id string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: Comment - comment"))
}
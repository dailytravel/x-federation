package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/dailytravel/x/workspace/graph/model"
)

// ID is the resolver for the id field.
func (r *boardResolver) ID(ctx context.Context, obj *model.Board) (string, error) {
	return obj.ID.Hex(), nil
}

// Portfolio is the resolver for the portfolio field.
func (r *boardResolver) Portfolio(ctx context.Context, obj *model.Board) (*model.Portfolio, error) {
	panic(fmt.Errorf("not implemented: Portfolio - portfolio"))
}

// DueDate is the resolver for the due_date field.
func (r *boardResolver) DueDate(ctx context.Context, obj *model.Board) (*string, error) {
	panic(fmt.Errorf("not implemented: DueDate - due_date"))
}

// Metadata is the resolver for the metadata field.
func (r *boardResolver) Metadata(ctx context.Context, obj *model.Board) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// Followers is the resolver for the followers field.
func (r *boardResolver) Followers(ctx context.Context, obj *model.Board) ([]*model.Follow, error) {
	panic(fmt.Errorf("not implemented: Followers - followers"))
}

// CreatedAt is the resolver for the created_at field.
func (r *boardResolver) CreatedAt(ctx context.Context, obj *model.Board) (string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - created_at"))
}

// UpdatedAt is the resolver for the updated_at field.
func (r *boardResolver) UpdatedAt(ctx context.Context, obj *model.Board) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// Lists is the resolver for the lists field.
func (r *boardResolver) Lists(ctx context.Context, obj *model.Board) ([]*model.List, error) {
	panic(fmt.Errorf("not implemented: Lists - lists"))
}

// CreateBoard is the resolver for the createBoard field.
func (r *mutationResolver) CreateBoard(ctx context.Context, input model.NewBoard) (*model.Board, error) {
	panic(fmt.Errorf("not implemented: CreateBoard - createBoard"))
}

// UpdateBoard is the resolver for the updateBoard field.
func (r *mutationResolver) UpdateBoard(ctx context.Context, id string, input model.UpdateBoard) (*model.Board, error) {
	panic(fmt.Errorf("not implemented: UpdateBoard - updateBoard"))
}

// DeleteBoard is the resolver for the deleteBoard field.
func (r *mutationResolver) DeleteBoard(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteBoard - deleteBoard"))
}

// DeleteBoards is the resolver for the deleteBoards field.
func (r *mutationResolver) DeleteBoards(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeleteBoards - deleteBoards"))
}

// Board is the resolver for the board field.
func (r *queryResolver) Board(ctx context.Context, id string) (*model.Board, error) {
	panic(fmt.Errorf("not implemented: Board - board"))
}

// Boards is the resolver for the boards field.
func (r *queryResolver) Boards(ctx context.Context, args map[string]interface{}) (*model.Boards, error) {
	panic(fmt.Errorf("not implemented: Boards - boards"))
}

// Board returns BoardResolver implementation.
func (r *Resolver) Board() BoardResolver { return &boardResolver{r} }

type boardResolver struct{ *Resolver }
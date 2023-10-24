package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/cms/graph/model"
	"github.com/dailytravel/x/cms/internal/utils"
	"github.com/dailytravel/x/cms/pkg/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	uid, err := utils.UID(ctx)
	if err != nil {
		return nil, err
	}

	item := &model.Post{
		UID:         *uid,
		Type:        input.Type,
		Locale:      input.Locale,
		Slug:        input.Slug,
		Commentable: *input.Commentable,
		Status:      *input.Status,
		Title:       map[string]interface{}{input.Locale: input.Title},
		Summary:     map[string]interface{}{input.Locale: input.Summary},
		Body:        map[string]interface{}{input.Locale: input.Body},
		Model: model.Model{
			Metadata: input.Metadata,
		},
	}

	if input.Parent != nil {
		parent, err := primitive.ObjectIDFromHex(*input.Parent)
		if err != nil {
			return nil, err
		}
		item.Parent = &parent
	}

	res, err := r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	item.ID = res.InsertedID.(primitive.ObjectID)

	return item, nil
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input model.UpdatePost) (*model.Post, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Fetch the existing category entry
	filter := bson.M{"_id": _id}
	item := &model.Post{}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	if input.Parent != nil {
		parent, err := primitive.ObjectIDFromHex(*input.Parent)
		if err != nil {
			return nil, err
		}
		item.Parent = &parent
	}

	// Update fields based on input
	if input.Title != nil {
		item.Title[input.Locale] = *input.Title
	}
	if input.Slug != nil {
		item.Slug = input.Slug
	}
	if input.Body != nil {
		item.Body[input.Locale] = *input.Body
	}
	if input.Metadata != nil {
		item.Metadata = input.Metadata
	}

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

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeletePost - deletePost"))
}

// DeletePosts is the resolver for the deletePosts field.
func (r *mutationResolver) DeletePosts(ctx context.Context, ids []string) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DeletePosts - deletePosts"))
}

// ID is the resolver for the id field.
func (r *postResolver) ID(ctx context.Context, obj *model.Post) (string, error) {
	return obj.ID.Hex(), nil
}

// Title is the resolver for the title field.
func (r *postResolver) Title(ctx context.Context, obj *model.Post) (string, error) {
	locale := auth.Locale(ctx)

	if locale == nil {
		locale = &obj.Locale
	}

	if title, ok := obj.Title[*locale].(string); ok {
		return title, nil
	}

	return obj.Title[obj.Locale].(string), nil
}

// Summary is the resolver for the summary field.
func (r *postResolver) Summary(ctx context.Context, obj *model.Post) (string, error) {
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	if summary, ok := obj.Summary[*locale].(string); ok {
		return summary, nil
	}

	return obj.Summary[obj.Locale].(string), nil
}

// Body is the resolver for the body field.
func (r *postResolver) Body(ctx context.Context, obj *model.Post) (string, error) {
	locale := auth.Locale(ctx)
	if locale == nil {
		locale = &obj.Locale
	}

	if body, ok := obj.Body[*locale].(string); ok {
		return body, nil
	}

	return obj.Body[obj.Locale].(string), nil
}

// Metadata is the resolver for the metadata field.
func (r *postResolver) Metadata(ctx context.Context, obj *model.Post) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// UID is the resolver for the uid field.
func (r *postResolver) UID(ctx context.Context, obj *model.Post) (string, error) {
	return obj.ID.Hex(), nil
}

// Created is the resolver for the created field.
func (r *postResolver) Created(ctx context.Context, obj *model.Post) (string, error) {
	return time.Unix(int64(obj.Created.T), 0).Format(time.RFC3339), nil
}

// Updated is the resolver for the updated field.
func (r *postResolver) Updated(ctx context.Context, obj *model.Post) (string, error) {
	return time.Unix(int64(obj.Updated.T), 0).Format(time.RFC3339), nil
}

// Terms is the resolver for the terms field.
func (r *postResolver) Terms(ctx context.Context, obj *model.Post) ([]*model.Term, error) {
	var items []*model.Term

	filter := bson.M{"_id": bson.M{"$in": obj.Terms}}
	options := options.Find().SetProjection(bson.M{"_id": 1, "name": 1, "slug": 1, "post": 1})

	cursor, err := r.db.Collection("terms").Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

// Parent is the resolver for the parent field.
func (r *postResolver) Parent(ctx context.Context, obj *model.Post) (*model.Post, error) {
	if obj.Parent == nil {
		return nil, nil
	}
	var item *model.Post

	filter := bson.M{"_id": obj.Parent}
	options := options.FindOne().SetProjection(bson.M{"_id": 1, "name": 1, "email": 1, "photos": 1})

	err := r.db.Collection(item.Collection()).FindOne(ctx, filter, options).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// Images is the resolver for the images field.
func (r *postResolver) Images(ctx context.Context, obj *model.Post) ([]*model.Image, error) {
	var items []*model.Image

	filter := bson.M{"object._id": obj.ID, "object.collection": obj.Collection()}

	cursor, err := r.db.Collection("images").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	var item *model.Post

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": _id}

	if err := r.db.Collection("posts").FindOne(ctx, filter).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, filter map[string]interface{}, project map[string]interface{}, sort map[string]interface{}, collation map[string]interface{}, limit *int, skip *int) (*model.Posts, error) {
	var items []*model.Post

	// Convert map to bson.M which is a type alias for map[string]interface{}
	_filter := utils.Filter(filter)
	opts := utils.Sort(sort)

	if project != nil {
		opts.SetProjection(project)
	}
	if limit != nil {
		opts.SetLimit(int64(*limit))
	}
	if skip != nil {
		opts.SetSkip(int64(*skip))
	}

	cursor, err := r.db.Collection("posts").Find(ctx, _filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	//get total count
	count, err := r.db.Collection("posts").CountDocuments(ctx, _filter, nil)
	if err != nil {
		return nil, err
	}

	return &model.Posts{
		Count: int(count),
		Data:  items,
	}, nil
}

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }

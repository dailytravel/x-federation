package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"time"

	"github.com/dailytravel/x/hrm/auth"
	"github.com/dailytravel/x/hrm/graph/model"
	"github.com/dailytravel/x/hrm/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateResume is the resolver for the createResume field.
func (r *mutationResolver) CreateResume(ctx context.Context, input model.CreateResume) (*model.Resume, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	item := &model.Resume{
		UID:            *uid,
		Title:          input.Title,
		Summary:        input.Summary,
		Experience:     input.Experience,
		Education:      input.Education,
		Skills:         input.Skills,
		Certifications: input.Certifications,
		Languages:      input.Languages,
		Projects:       input.Projects,
		References:     input.References,
		Model: model.Model{
			CreatedBy: uid,
			UpdatedBy: uid,
		},
	}

	_, err = r.db.Collection(item.Collection()).InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateResume is the resolver for the updateResume field.
func (r *mutationResolver) UpdateResume(ctx context.Context, id string, input model.UpdateResume) (*model.Resume, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find the item by ID
	item := &model.Resume{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	// Update the resume fields with the provided input
	if input.Title != nil {
		item.Title = *input.Title
	}

	if input.Summary != nil {
		item.Summary = *input.Summary
	}

	if input.Experience != nil {
		item.Experience = *input.Experience
	}

	if input.Education != nil {
		item.Education = *input.Education
	}

	if input.Skills != nil {
		item.Skills = *input.Skills
	}

	if input.Certifications != nil {
		item.Certifications = *input.Certifications
	}

	if input.Languages != nil {
		item.Languages = *input.Languages
	}

	if input.Projects != nil {
		item.Projects = *input.Projects
	}

	if input.References != nil {
		item.References = *input.References
	}

	// Update the "updated_by" field
	item.UpdatedBy = uid

	// Perform the update operation in the database
	_, err = r.db.Collection(item.Collection()).UpdateOne(ctx, filter, bson.M{"$set": item})
	if err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteResume is the resolver for the deleteResume field.
func (r *mutationResolver) DeleteResume(ctx context.Context, id string) (map[string]interface{}, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find the resume record by ID
	item := &model.Resume{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	// Update the resume record to mark it as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"deleted_by": uid,
			"status":     "deleted",
			"updated_by": uid,
			"updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Perform the update operation in the database
	var deletedResume model.Resume
	err = r.db.Collection(item.Collection()).FindOneAndUpdate(ctx, filter, update, opts).Decode(&deletedResume)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "deletedCount": 1}, nil
}

// DeleteResumes is the resolver for the deleteResumes field.
func (r *mutationResolver) DeleteResumes(ctx context.Context, ids []string) (map[string]interface{}, error) {
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
	result, err := r.db.Collection("resumes").UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "deletedCount": result.ModifiedCount}, nil
}

// Resume is the resolver for the resume field.
func (r *queryResolver) Resume(ctx context.Context, id string) (*model.Resume, error) {
	var item *model.Resume
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Collection("employees").FindOne(ctx, bson.M{"_id": _id}).Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}

// Resumes is the resolver for the resumes field.
func (r *queryResolver) Resumes(ctx context.Context, args map[string]interface{}) (*model.Resumes, error) {
	var items []*model.Resume
	//find all items
	cur, err := r.db.Collection("resumes").Find(ctx, utils.Query(args), utils.Options(args))
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Resume
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("resumes").CountDocuments(ctx, utils.Query(args), nil)
	if err != nil {
		return nil, err
	}

	return &model.Resumes{
		Count: int(count),
		Data:  items,
	}, nil
}

// ID is the resolver for the id field.
func (r *resumeResolver) ID(ctx context.Context, obj *model.Resume) (string, error) {
	return obj.ID.Hex(), nil
}

// Metadata is the resolver for the metadata field.
func (r *resumeResolver) Metadata(ctx context.Context, obj *model.Resume) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *resumeResolver) CreatedAt(ctx context.Context, obj *model.Resume) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *resumeResolver) UpdatedAt(ctx context.Context, obj *model.Resume) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// UID is the resolver for the uid field.
func (r *resumeResolver) UID(ctx context.Context, obj *model.Resume) (string, error) {
	return obj.UID.Hex(), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *resumeResolver) CreatedBy(ctx context.Context, obj *model.Resume) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	createdBy := obj.CreatedBy.Hex()

	return &createdBy, nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *resumeResolver) UpdatedBy(ctx context.Context, obj *model.Resume) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	updatedBy := obj.UpdatedBy.Hex()

	return &updatedBy, nil
}

// Resume returns ResumeResolver implementation.
func (r *Resolver) Resume() ResumeResolver { return &resumeResolver{r} }

type resumeResolver struct{ *Resolver }

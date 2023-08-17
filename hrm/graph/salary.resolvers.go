package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dailytravel/x/hrm/auth"
	"github.com/dailytravel/x/hrm/graph/model"
	"github.com/dailytravel/x/hrm/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateSalary is the resolver for the createSalary field.
func (r *mutationResolver) CreateSalary(ctx context.Context, input model.NewSalary) (*model.Salary, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	item := &model.Salary{
		Amount:   input.Amount,
		Currency: input.Currency,
		Model: model.Model{
			CreatedBy: uid,
			UpdatedBy: uid,
		},
	}

	// Insert the new organization
	if _, err := r.db.Collection(item.Collection()).InsertOne(ctx, item, nil); err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateSalary is the resolver for the updateSalary field.
func (r *mutationResolver) UpdateSalary(ctx context.Context, id string, input model.UpdateSalary) (*model.Salary, error) {
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
	item := &model.Salary{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	item.UpdatedBy = uid

	// Update the position in the database
	if err := r.db.Collection(item.Collection()).FindOneAndUpdate(ctx, filter, item).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

// DeleteSalary is the resolver for the deleteSalary field.
func (r *mutationResolver) DeleteSalary(ctx context.Context, id string) (map[string]interface{}, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	// Convert the ID string to ObjectID
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find the salary by ID
	salary := &model.Salary{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(salary.Collection()).FindOne(ctx, filter).Decode(salary)
	if err != nil {
		return nil, err
	}

	// Define the update to mark the salary as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"deleted_by": uid,
			"status":     "deleted",
			"updated_by": uid,
			"updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Update the salary in the database
	if _, err := r.db.Collection(salary.Collection()).UpdateOne(ctx, filter, update); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success"}, nil
}

// DeleteSalaries is the resolver for the deleteSalaries field.
func (r *mutationResolver) DeleteSalaries(ctx context.Context, ids []string) (map[string]interface{}, error) {
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

	// Define the update to mark salaries as deleted
	update := bson.M{
		"$set": bson.M{
			"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
			"deleted_by": uid,
			"status":     "deleted",
			"updated_by": uid,
			"updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	// Update the salaries in the database
	result, err := r.db.Collection("salaries").UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "deletedCount": result.ModifiedCount}, nil
}

// Salary is the resolver for the salary field.
func (r *queryResolver) Salary(ctx context.Context, id string) (*model.Salary, error) {
	var item *model.Salary

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": _id}

	if err := r.db.Collection("salaries").FindOne(ctx, filter).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// Salaries is the resolver for the salaries field.
func (r *queryResolver) Salaries(ctx context.Context, args map[string]interface{}) (*model.Salaries, error) {
	var items []*model.Salary
	//find all items
	cur, err := r.db.Collection("salaries").Find(ctx, utils.Query(args), utils.Options(args))
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Salary
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("salaries").CountDocuments(ctx, utils.Query(args), nil)
	if err != nil {
		return nil, err
	}

	return &model.Salaries{
		Count: int(count),
		Data:  items,
	}, nil
}

// ID is the resolver for the id field.
func (r *salaryResolver) ID(ctx context.Context, obj *model.Salary) (string, error) {
	return obj.ID.Hex(), nil
}

// StartDate is the resolver for the start_date field.
func (r *salaryResolver) StartDate(ctx context.Context, obj *model.Salary) (string, error) {
	return obj.StartDate.Time().String(), nil
}

// EndDate is the resolver for the end_date field.
func (r *salaryResolver) EndDate(ctx context.Context, obj *model.Salary) (*string, error) {
	endDate := obj.EndDate.Time().String()
	return &endDate, nil
}

// Metadata is the resolver for the metadata field.
func (r *salaryResolver) Metadata(ctx context.Context, obj *model.Salary) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *salaryResolver) CreatedAt(ctx context.Context, obj *model.Salary) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *salaryResolver) UpdatedAt(ctx context.Context, obj *model.Salary) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// Employee is the resolver for the employee field.
func (r *salaryResolver) Employee(ctx context.Context, obj *model.Salary) (*model.Employee, error) {
	var item *model.Employee

	filter := bson.M{"_id": obj.Employee}
	if err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("employee not found")
		}
		return nil, err
	}

	return item, nil
}

// CreatedBy is the resolver for the created_by field.
func (r *salaryResolver) CreatedBy(ctx context.Context, obj *model.Salary) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	createdBy := obj.CreatedBy.Hex()

	return &createdBy, nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *salaryResolver) UpdatedBy(ctx context.Context, obj *model.Salary) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	updatedBy := obj.UpdatedBy.Hex()

	return &updatedBy, nil
}

// Salary returns SalaryResolver implementation.
func (r *Resolver) Salary() SalaryResolver { return &salaryResolver{r} }

type salaryResolver struct{ *Resolver }

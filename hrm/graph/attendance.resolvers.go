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
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ID is the resolver for the id field.
func (r *attendanceResolver) ID(ctx context.Context, obj *model.Attendance) (string, error) {
	return obj.ID.Hex(), nil
}

// Employee is the resolver for the employee field.
func (r *attendanceResolver) Employee(ctx context.Context, obj *model.Attendance) (*model.Employee, error) {
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

// TimeIn is the resolver for the time_in field.
func (r *attendanceResolver) TimeIn(ctx context.Context, obj *model.Attendance) (string, error) {
	return time.Unix(int64(obj.TimeIn.T), 0).Format(time.RFC3339), nil
}

// TimeOut is the resolver for the time_out field.
func (r *attendanceResolver) TimeOut(ctx context.Context, obj *model.Attendance) (string, error) {
	return time.Unix(int64(obj.TimeOut.T), 0).Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *attendanceResolver) Metadata(ctx context.Context, obj *model.Attendance) (map[string]interface{}, error) {
	return obj.Metadata, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *attendanceResolver) CreatedAt(ctx context.Context, obj *model.Attendance) (string, error) {
	return time.Unix(int64(obj.CreatedAt.T), 0).Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *attendanceResolver) UpdatedAt(ctx context.Context, obj *model.Attendance) (string, error) {
	return time.Unix(int64(obj.UpdatedAt.T), 0).Format(time.RFC3339), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *attendanceResolver) CreatedBy(ctx context.Context, obj *model.Attendance) (*string, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	createdBy := obj.CreatedBy.Hex()

	return &createdBy, nil
}

// UpdatedBy is the resolver for the updated_by field.
func (r *attendanceResolver) UpdatedBy(ctx context.Context, obj *model.Attendance) (*string, error) {
	if obj.UpdatedBy == nil {
		return nil, nil
	}

	updatedBy := obj.UpdatedBy.Hex()

	return &updatedBy, nil
}

// CheckIn is the resolver for the checkIn field.
func (r *mutationResolver) CheckIn(ctx context.Context, employee string) (*model.Attendance, error) {
	_employee, err := primitive.ObjectIDFromHex(employee)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"employee": _employee,
		"time_in": bson.M{
			"$gte": primitive.Timestamp{T: uint32(time.Now().Unix()) - 86400},
			"$lt":  primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
		"time_out": bson.M{
			"$exists": false,
		},
	}

	var item *model.Attendance
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			item = &model.Attendance{
				Employee: _employee,
				TimeIn:   primitive.Timestamp{T: uint32(time.Now().Unix())},
			}

			res, err := r.db.Collection(item.Collection()).InsertOne(ctx, item, nil)
			if err != nil {
				return nil, err
			}

			item.ID = res.InsertedID.(primitive.ObjectID)

			return item, nil
		}

		return nil, err
	}

	return item, nil
}

// CheckOut is the resolver for the checkOut field.
func (r *mutationResolver) CheckOut(ctx context.Context, employee string) (*model.Attendance, error) {
	_employee, err := primitive.ObjectIDFromHex(employee)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"employee": _employee,
		"time_in": bson.M{
			"$gte": primitive.Timestamp{T: uint32(time.Now().Unix()) - 86400},
			"$lt":  primitive.Timestamp{T: uint32(time.Now().Unix())},
		},
	}

	var item model.Attendance

	err = r.db.Collection(item.Collection()).FindOne(ctx, filter, options.FindOne().SetSort(bson.D{{Key: "time_in", Value: -1}})).Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("attendance not found")
		}
		return nil, err
	}

	out := primitive.Timestamp{T: uint32(time.Now().Unix())}
	update := &model.Attendance{
		TimeOut: &out,
	}

	_, err = r.db.Collection(item.Collection()).UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// CreateAttendance is the resolver for the createAttendance field.
func (r *mutationResolver) CreateAttendance(ctx context.Context, input model.NewAttendance) (*model.Attendance, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	// Create a new attendance
	item := &model.Attendance{
		Status: input.Status,
		Model: model.Model{
			CreatedBy: uid,
			UpdatedBy: uid,
			Metadata:  input.Metadata,
		},
	}

	// Perform the insertion into the database
	_, err = r.db.Collection("attendances").InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateAttendance is the resolver for the updateAttendance field.
func (r *mutationResolver) UpdateAttendance(ctx context.Context, id string, input model.UpdateAttendance) (*model.Attendance, error) {
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
	item := &model.Attendance{}
	filter := bson.M{"_id": _id}
	err = r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	// Update fields if provided in input
	if input.Status != nil {
		item.Status = *input.Status
	}
	if input.Notes != nil {
		item.Notes = *input.Notes
	}

	// Update the attendance record in the database
	update := bson.M{
		"$set": bson.M{
			"status":     item.Status,
			"notes":      item.Notes,
			"updated_by": uid,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updated model.Attendance
	err = r.db.Collection("attendances").FindOneAndUpdate(ctx, filter, update, opts).Decode(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

// DeleteAttendance is the resolver for the deleteAttendance field.
func (r *mutationResolver) DeleteAttendance(ctx context.Context, id string) (map[string]interface{}, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	// Convert the ID string to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Create a filter to match the provided ID
	filter := bson.M{"_id": objectID}

	// Update the attendance record to mark it as deleted
	archiveFields := bson.M{
		"status":     "archived",
		"updated_by": uid,
		"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
	}

	update := bson.M{"$set": archiveFields}
	updateResult, err := r.db.Collection("attendances").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if updateResult.MatchedCount == 0 {
		return nil, fmt.Errorf("attendance not found")
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}

// DeleteAttendances is the resolver for the deleteAttendances field.
func (r *mutationResolver) DeleteAttendances(ctx context.Context, ids []string) (map[string]interface{}, error) {
	uid, err := utils.UID(auth.Auth(ctx))
	if err != nil {
		return nil, err
	}

	// Convert string IDs to ObjectIDs
	objectIDs := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs[i] = objectID
	}

	// Create a filter to match the provided IDs
	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	// Update the attendance records to mark them as deleted
	archiveFields := bson.M{
		"status":     "archived",
		"updated_by": uid,
		"deleted_at": primitive.Timestamp{T: uint32(time.Now().Unix())},
	}

	update := bson.M{"$set": archiveFields}
	updateResult, err := r.db.Collection("attendances").UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"success":      true,
		"deletedCount": updateResult.ModifiedCount,
	}, nil
}

// Attendance is the resolver for the attendance field.
func (r *queryResolver) Attendance(ctx context.Context, id string) (*model.Attendance, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var item *model.Attendance
	filter := bson.M{"_id": _id}

	if err := r.db.Collection(item.Collection()).FindOne(ctx, filter).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found for filter %v", filter)
		}
		return nil, err
	}

	return item, nil
}

// Attendances is the resolver for the attendances field.
func (r *queryResolver) Attendances(ctx context.Context, args map[string]interface{}) (*model.Attendances, error) {
	var items []*model.Attendance

	filter := utils.Query(args).(bson.M)

	opts := utils.Options(args)
	opts.SetSort(bson.M{"time_in": -1})

	cur, err := r.db.Collection("attendances").Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item *model.Attendance
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	//get total count
	count, err := r.db.Collection("attendances").CountDocuments(ctx, filter, nil)
	if err != nil {
		return nil, err
	}

	return &model.Attendances{
		Count: int(count),
		Data:  items,
	}, nil
}

// Attendance returns AttendanceResolver implementation.
func (r *Resolver) Attendance() AttendanceResolver { return &attendanceResolver{r} }

type attendanceResolver struct{ *Resolver }

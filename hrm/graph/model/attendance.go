package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Attendance struct {
	Model   `bson:",inline"`
	Owner   primitive.ObjectID   `bson:"owner" json:"owner"`
	TimeIn  primitive.Timestamp  `json:"time_in" bson:"time_in"`
	TimeOut *primitive.Timestamp `json:"time_out" bson:"time_out"`
	Notes   string               `json:"notes,omitempty" bson:"notes,omitempty"`
	Status  string               `json:"status" bson:"status"`
}

func (i *Attendance) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Attendance
	return bson.Marshal((*t)(i))
}

func (i *Attendance) Collection() string {
	return "attendances"
}

func (i *Attendance) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "owner", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "time_in", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "time_out", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
	}
}
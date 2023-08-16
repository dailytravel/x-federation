package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Board struct {
	Model        `bson:",inline"`
	UID          primitive.ObjectID  `json:"uid" bson:"uid"`
	Organization *primitive.ObjectID `json:"organization,omitempty" bson:"organization,omitempty"`
	Portfolio    *primitive.ObjectID `json:"portfolio,omitempty" bson:"portfolio,omitempty"`
	Type         string              `json:"type" bson:"type"`
	Title        string              `json:"title" bson:"title"`
	Description  *string             `json:"description,omitempty" bson:"description,omitempty"`
	IsTemplate   bool                `json:"is_template,omitempty" bson:"is_template,omitempty"`
	DueDate      primitive.Timestamp `json:"due_date,omitempty" bson:"due_date,omitempty"`
	Starred      bool                `json:"starred,omitempty" bson:"starred,omitempty"`
	Background   string              `json:"background,omitempty" bson:"background,omitempty"`
	Order        int                 `json:"order" bson:"order"`
	Status       string              `json:"status" bson:"status"`
}

func (Board) IsEntity() {}

func (i *Board) MarshalBSON() ([]byte, error) {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}

	if i.CreatedAt.IsZero() {
		i.CreatedAt = now
	}

	i.UpdatedAt = now

	type t Board
	return bson.Marshal((*t)(i))
}

func (i *Board) Collection() string {
	return "boards"
}

func (i *Board) Index() []mongo.IndexModel {
	return []mongo.IndexModel{
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "organization", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "type", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "starred", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "order", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "status", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "is_template", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "due_date", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index()},
		{Keys: bson.D{{Key: "deleted_at", Value: 1}}, Options: options.Index()},
	}
}

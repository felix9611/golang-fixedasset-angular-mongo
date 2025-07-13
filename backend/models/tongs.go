package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tongs struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Status int                `bson:"status" json:"status"` // 0: Available, 1: In Use
}
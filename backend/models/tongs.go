package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tongs struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Status int                `bson:"status" json:"status"` // 0: Available, 1: In Use
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
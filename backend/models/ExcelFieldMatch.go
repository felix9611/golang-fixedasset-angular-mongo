package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ExcelFieldMatchs struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FunctionCode string             `bson:"functionCode" json:"functionCode"`
	FunctionName string             `bson:"functionName" json:"functionName"`
	FunctionType string             `bson:"functionType" json:"functionType"`
	FieldLists   []any              `bson:"fieldLists" json:"fieldLists"`
	Status       int                `bson:"status" json:"status"`
	CreatedAt    time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}
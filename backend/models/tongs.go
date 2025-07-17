package models
/*
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
*/
type Tongs struct {
	Base
	Name  string             `bson:"name" json:"name"`
}
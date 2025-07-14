package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tongs struct {
	Base
	Name  string             `bson:"name" json:"name"`
	Status int                `bson:"status" json:"status"`
}
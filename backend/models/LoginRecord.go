package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LoginRecords struct {
	ID        	primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	username  	string             `bson:"username" json:"username"`
	ipAddress 	string             `bson:"ipAddress" json:"ipAddress"`
	loginStatus string            	`bson:"loginStatus" json:"loginStatus"`
	loginTime 	time.Time          `bson:"loginTime" json:"loginTime"`
}
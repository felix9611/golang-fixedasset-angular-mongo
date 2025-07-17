package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SysRoles struct {
	Base
	Name        string             `bson:"name" json:"name"`
	Code 	  string             `bson:"code" json:"code"`
	Remark	  string             `bson:"remark" json:"remark"`
	MenuIds []primitive.ObjectID	`bson:"menuIds,omitempty" json:"menuIds"`
	Read  bool               `bson:"read" json:"read"`
	Write bool               `bson:"write" json:"write"`
	Delete bool             `bson:"delete" json:"delete"`
	Upload bool             `bson:"upload" json:"upload"`
	Update bool             `bson:"update" json:"update"`
}
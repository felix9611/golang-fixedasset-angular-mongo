package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SysUsers struct {
	Base
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Email       string             `bson:"email" json:"email"`
	AvatarBase64 string             `bson:"avatarBase64" json:"avatarBase64"`
	DeptId primitive.ObjectID `bson:"DeptId,omitempty"`
	roles []primitive.ObjectID `bson:"roles,omitempty" json:"roles"`
}
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Department struct {
	Base
	DeptCode string             `bson:"deptCode" json:"deptCode"`
	DeptName string             `bson:"deptName" json:"deptName"`
	Remark   string             `bson:"remark" json:"remark"`
}

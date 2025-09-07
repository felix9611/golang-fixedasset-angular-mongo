package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Budgets struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	DeptID   primitive.ObjectID `bson:"deptId" json:"deptId"`
	PlaceId  primitive.ObjectID `bson:"placeId" json:"placeId"`
	BudgetNo   string           `bson:"budgetNo" json:"budgetNo"`
	BudgetName string           `bson:"budgetName" json:"budgetName"`
	Year       string           `bson:"year" json:"year"`
	Month      string           `bson:"month" json:"month"`
	BudgetAmount float64       `bson:"budgetAmount" json:"budgetAmount"`
	BudgetFrom string           `bson:"budgetFrom" json:"budgetFrom"`
	BudgetTo   string           `bson:"budgetTo" json:"budgetTo"`
	BudgetStatus string         `bson:"budgetStatus" json:"budgetStatus"`
	Remark     string           `bson:"remark" json:"remark"`
	Status    int                `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}
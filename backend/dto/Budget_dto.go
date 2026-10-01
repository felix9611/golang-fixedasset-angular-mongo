package dto

import (
	"golang-fixedasset-mongo-backend/backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListBudgetRecordsDto struct {
	Name    string `json:"name"`
	Page    int64  `json:"page"`
	DeptID  string `json:"deptId"`
	PlaceId string `json:"placeId"`
	// Date    DateRange `json:"date"`
	Limit int64 `json:"limit"`
}

type DateRange struct {
	From primitive.DateTime `json:"from"`
	To   primitive.DateTime `json:"to"`
}

type BudgetSummary struct {
	BudgetAmount float64 `bson:"budgetAmount" json:"budgetAmount"`
	YearMonth    string  `bson:"yearMonth" json:"yearMonth"`
	Year         int     `bson:"year" json:"year"`
	Month        int     `bson:"month" json:"month"`
}

type BudgetList struct {
	Lists []models.Budgets `json:"lists"`
	Total int64            `json:"total"`
	Page  int              `json:"page"`
	Limit int              `json:"limit"`
}

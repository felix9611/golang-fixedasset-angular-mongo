package dto

import "golang-fixedasset-mongo-backend/backend/models"

type ListBudgetRecordsDto struct {
	Name  string `json:"name"`
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
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

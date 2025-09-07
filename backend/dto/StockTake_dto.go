package dto

type ListStockTakeDto struct {
	PlaceIds []string `json:"placeIds"`
	Name     string   `json:"name"`
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
}
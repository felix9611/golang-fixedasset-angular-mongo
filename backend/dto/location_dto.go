package dto

type LocationPageDto struct {
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
	Name string `json:"name"`
}
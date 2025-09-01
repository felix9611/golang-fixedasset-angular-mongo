package dto

type VendorPageDto struct {
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
	Name  string `json:"name"`
	Place string `json:"place"`
	Contact string `json:"contact"`
}
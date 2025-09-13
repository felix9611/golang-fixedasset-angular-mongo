package dto

type ExcelFieldMatchPageDto struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
}

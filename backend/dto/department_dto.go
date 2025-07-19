package dto

type DepartmentPageDto struct {
	Name  string `json:"name"`
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
}

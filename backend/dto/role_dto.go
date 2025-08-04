package dto

type DepartmentPageDto struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
}

package dto

type CodeTypeListDto struct {
	Name  string `json:"name"`
	Limit int64  `json:"limit"`
	Page  int64  `json:"page"`
}

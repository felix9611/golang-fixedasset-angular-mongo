package dto

type RolesPageDto struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
}

type MenuItemPermissionBody struct {
	MenuIds []any `json:"menuIds"`
	ID      string `json:"id"`
}

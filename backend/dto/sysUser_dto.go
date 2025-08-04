package dto

type SysUserPageDto struct {
	RoleIds []string `json:"roleIds"`
	Name    string   `json:"name"`
	Page    int64    `json:"page"`
	Limit   int64    `json:"limit"`
}

type SysUserUpdateDto struct {
	Username string `json:"username"`
	NewPassword string `json:"newPassword"`
}
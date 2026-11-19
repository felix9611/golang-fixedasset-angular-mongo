package dto

import (
	"golang-fixedasset-mongo-backend/backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserBody struct {
	Key      string          `json:"key"`
	UserData models.SysUsers `json:"userData"`
}

type SysUserPageDto struct {
	RoleIds  []primitive.ObjectID `json:"roleIds"`
	PlaceIds []primitive.ObjectID `json:"placeIds"`
	DeptIds  []primitive.ObjectID `json:"deptIds"`
	Name     string               `json:"name"`
	Page     int64                `json:"page"`
	Limit    int64                `json:"limit"`
}

type SysUserList struct {
	Lists []models.SysUsers `json:"lists"`
	Total int64             `json:"total"`
	Page  int               `json:"page"`
	Limit int               `json:"limit"`
}

type SysUserUpdatePasswordDto struct {
	Username    string `json:"username"`
	NewPassword string `json:"newPassword"`
}

type SysUserAvatarUpdateDto struct {
	Username  string `json:"username"`
	PhotoBase string `json:"photo"`
}

type SysUserCreateResponseDto struct {
	ID string `json:"id"`
}

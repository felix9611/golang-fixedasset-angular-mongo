package controllers

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"golang-fixedasset-mongo-backend/backend/dto"
	"net/http"
//	"time"

	"github.com/gin-gonic/gin"

)


func CreateSysUser(c *gin.Context) {
	var user models.SysUsers
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.CreateSysUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func UpdateSysUserByID(c *gin.Context) {
	var user models.SysUsers
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateSysUserByID(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func GetSysUsersById(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetOneSysUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func InactiveUserByID(c *gin.Context) {
	id := c.Param("id")
	result, err := services.InactiveUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to inactive user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func SysUserLists(c *gin.Context) {
	var pageDto dto.SysUserPageDto
	if err := c.ShouldBindJSON(&pageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.SysUserList(&pageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func SysUserAvatarUpdate(c *gin.Context) {
	var avatarDto dto.SysUserAvatarUpdateDto
	if err := c.ShouldBindJSON(&avatarDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UserUpdateAvatar(avatarDto.Username, avatarDto.PhotoBase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update avatar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func RegisterSysUserRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/sys/user")
	{
		userGroup.POST("/create", CreateSysUser)
		userGroup.POST("/update", UpdateSysUserByID)
		userGroup.GET("/one/:id", GetSysUsersById)
		userGroup.DELETE("/:id", InactiveUserByID)
		userGroup.POST("/list", SysUserLists)
		userGroup.POST("/user-self/update-avatar", SysUserAvatarUpdate)
	}
}

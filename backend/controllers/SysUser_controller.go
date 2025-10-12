package controllers

import (
	"golang-fixedasset-mongo-backend/backend/dto"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"net/http"

	//	"time"
	jwt "github.com/appleboy/gin-jwt/v2"
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

// @Summary      Get one Sys User record by id
// @Description  Get one Sys User record by id
// @Tags         Sys User
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Sys User ID"
// @Success      200      {object}  models.SysUsers
// @Router       /sys/user/one/{id} [get]
func GetSysUsersById(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetOneSysUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	c.JSON(http.StatusOK, user)
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

func SysUserPasswordUpdate(c *gin.Context) {
	var passwordDto dto.SysUserUpdatePasswordDto
	if err := c.ShouldBindJSON(&passwordDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateUserPassword(passwordDto.Username, passwordDto.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func RegisterSysUserRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	userGroup := rg.Group("/sys/user", handle.MiddlewareFunc())
	{
		userGroup.POST("/create", CreateSysUser)
		userGroup.POST("/update", UpdateSysUserByID)
		userGroup.GET("/one/:id", GetSysUsersById)
		userGroup.DELETE("/:id", InactiveUserByID)
		userGroup.POST("/list", SysUserLists)
		userGroup.POST("/user-self/update-avatar", SysUserAvatarUpdate)
		userGroup.POST("/user-self/update-password", SysUserPasswordUpdate)
	}
}

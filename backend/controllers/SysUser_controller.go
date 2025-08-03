package controllers

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	//"golang-fixedasset-mongo-backend/backend/dto"
	"net/http"
//	"time"

	"github.com/gin-gonic/gin"

)

var u models.SysUsers

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
	id := c.Param("id")
	var user models.SysUsers
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateSysUserByID(id, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func GetSysUsersById(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetOneSysUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func InactiveUserByID(c *gin.Context) {
	id := c.Param("id")
	result, err := services.InactiveSysUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to inactive user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}


func RegisterSysUserRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/sys-users")
	{
		userGroup.POST("/create", CreateSysUser)
		userGroup.POST("/:id", UpdateSysUserByID)
		userGroup.GET("/:id", GetSysUsersById)
		userGroup.DELETE("/:id", InactiveUserByID)
	}
}

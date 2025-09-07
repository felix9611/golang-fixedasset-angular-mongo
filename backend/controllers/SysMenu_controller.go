package controllers

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"golang-fixedasset-mongo-backend/backend/dto"
	"net/http"
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
)

func CreateSysMenu(c *gin.Context) {
	var data models.SysMenus
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.CreateSysMenu(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func ListAllSysMenu(c *gin.Context) {
	var query dto.SysMenuList
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.ListAllMenu(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu items"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetSysMenuMainIdList(c *gin.Context) {
	result, err := services.ListAllMainIdMenu()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu main IDs"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetMenuItemById(c *gin.Context) {
	id := c.Param("id")
	result, err := services.GetOneMenuItemById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdateMenuItemById(c *gin.Context) {
	var data models.SysMenus
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.UpdateMenuItem(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetAllMenuItems(c *gin.Context) {
	result, err := services.GetAllMenuItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu items"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func RegisterSysMenuRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	sysMenuRoute := rg.Group("/sys/menu", handle.MiddlewareFunc())
	{
		sysMenuRoute.POST("/create", CreateSysMenu)
		sysMenuRoute.POST("/list", ListAllSysMenu)
		sysMenuRoute.GET("/main-item", GetSysMenuMainIdList)
		sysMenuRoute.GET("/one/:id", GetMenuItemById)
		sysMenuRoute.POST("/update", UpdateMenuItemById)
		sysMenuRoute.GET("/all-menu", GetAllMenuItems)
	}
}
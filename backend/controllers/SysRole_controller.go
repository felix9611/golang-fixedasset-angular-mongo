package controllers

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateSysUser(c *gin.Context) {
	var role models.SysRole
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.CreateSysRole(&role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func GetRoleById(c *gin.Context) {
	id := c.Param("id")
	role, err := services.GetOneSysRoleById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": role})
}

func UpdateSysRole(c *gin.Context) {
	id := c.Param("id")
	var role models.SysRole
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateSysRoleById(id, &role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role updated successfully", "id": id, "data": result})
}

func InactiveSysRoleByID(c *gin.Context) {
	id := c.Param("id")
	result, err := services.InactiveSysRoleByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deactivate role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role deactivated successfully", "id": id, "data": result})
}

func ListPageRoles(c *gin.Context) {
	var rolePageDto dto.RolePageDto
	if err := c.ShouldBindJSON(&rolePageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.RolesList(&rolePageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list roles"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}	

func RegisterSysRoleRoutes(rg *gin.RouterGroup) {
	roleGroup := rg.Group("/sys-role")
	{
		roleGroup.POST("/create", CreateSysRole)
		roleGroup.GET("/one/:id", GetRoleById)
		roleGroup.POST("/update/:id", UpdateSysRole)
		roleGroup.DELETE("/void/:id", InactiveSysRoleByID)
		roleGroup.GET("/list", ListPageRoles)
	}
}

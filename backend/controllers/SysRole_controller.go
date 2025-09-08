package controllers

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"golang-fixedasset-mongo-backend/backend/dto"
	"net/http"
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
)

func CreateSysRoleApi(c *gin.Context) {
	var role models.SysRoles
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

func GetAllRoles(c *gin.Context) {
	roles, err := services.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get roles"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roles})
}

func UpdateSysRole(c *gin.Context) {
	// id := c.Param("id")
	var role models.SysRoles
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateRoleById(&role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role updated successfully", "id": role.ID, "data": result})
}

func InactiveSysRoleByID(c *gin.Context) {
	id := c.Param("id")
	result, err := services.VoidRoleById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deactivate role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role deactivated successfully", "id": id, "data": result})
}

func ListPageRoles(c *gin.Context) {
	var rolePageDto dto.RolesPageDto
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

func HandleMeunPermission(c *gin.Context) {
	var menuPermissionDto dto.MenuItemPermissionBody
	if err := c.ShouldBindJSON(&menuPermissionDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.HandleMenuPermission(&menuPermissionDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update menu permissions"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu permissions updated successfully", "data": result})
}

func LoadRoleWithMenu(c *gin.Context) {
	var roleIdsDto dto.RoleIdsBody
	if err := c.ShouldBindJSON(&roleIdsDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.LoadRoleWithMenu(&roleIdsDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load roles with menus"})
		return
	}
	c.JSON(http.StatusOK, result)
}


func RegisterSysRoleRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	roleGroup := rg.Group("/sys/role", handle.MiddlewareFunc())
	{
		roleGroup.POST("/create", CreateSysRoleApi)
		roleGroup.GET("/one/:id", GetRoleById)
		roleGroup.POST("/update", UpdateSysRole)
		roleGroup.DELETE("/void/:id", InactiveSysRoleByID)
		roleGroup.POST("/list", ListPageRoles)
		roleGroup.POST("/update-permission", HandleMeunPermission)
		roleGroup.GET("/all", GetAllRoles)
		roleGroup.POST("/list-permission", LoadRoleWithMenu)
	}
}

package controllers

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"golang-fixedasset-mongo-backend/backend/dto"
	"net/http"
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
)

func CreateCodeType(c *gin.Context) {
	var codeType models.CodeTypes
	if err := c.ShouldBindJSON(&codeType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreateCodeType(&codeType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetCodeTypeById(c *gin.Context) {
	id := c.Param("id")
	result, err := services.GetOneCodeType(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func InactiveCodeTypeByID(c *gin.Context) {
	id := c.Param("id")
	result,err := services.VoidOneCodeType(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "CodeType inactivated successfully", "data": result })
}

func ListCodeType(c *gin.Context) {
	var codeTypePageDto dto.CodeTypeListDto
	if err := c.ShouldBindJSON(&codeTypePageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CodeTypeList(&codeTypePageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateCodeTypeById(c *gin.Context) {
	var codeType models.CodeTypes
	if err := c.ShouldBindJSON(&codeType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateCodeType(&codeType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ListCodeTypeByType(c *gin.Context) {
	typeString := c.Param("type")

	result, err := services.ListCodeTypeByType(typeString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func RegisterCodeTypeRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	codeTypeGroup := rg.Group("/base/code-type", handle.MiddlewareFunc())
	{
		codeTypeGroup.POST("/create", CreateCodeType)
		codeTypeGroup.GET("/one/:id", GetCodeTypeById)
		codeTypeGroup.DELETE("/void/:id", InactiveCodeTypeByID)
		codeTypeGroup.POST("/list", ListCodeType)
		codeTypeGroup.POST("/update", UpdateCodeTypeById)
		codeTypeGroup.GET("/get-type/:type", ListCodeTypeByType)
	}
}
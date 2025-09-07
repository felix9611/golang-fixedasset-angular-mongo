package controllers

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"golang-fixedasset-mongo-backend/backend/dto"
	"net/http"
)

func CreateAssetType(c *gin.Context) {
	var assetType models.AssetTypes
	if err := c.ShouldBindJSON(&assetType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreateAssetType(&assetType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetAssetTypeById(c *gin.Context) {
	id := c.Param("id")
	result, err := services.GetOneAssetType(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func InactiveAssetType(c *gin.Context) {
	id := c.Param("id")
	result, err := services.VoidOneAssetType(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Asset Type inactivated successfully", "data": result })
}

func UpdateAssetType(c *gin.Context) {
	var assetType models.AssetTypes
	if err := c.ShouldBindJSON(&assetType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateAssetType(&assetType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetAssetTypes(c *gin.Context) {
	assetTypes, err := services.ListAllAssetType()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get asset types"})
		return
	}
	c.JSON(http.StatusOK, assetTypes)
}

func ListPageAssetTypes(c *gin.Context) {
	var assetTypePageDto dto.AssetTypeListDto
	if err := c.ShouldBindJSON(&assetTypePageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	assetTypes, err := services.ListAssetType(&assetTypePageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list asset types"})
		return
	}
	c.JSON(http.StatusOK, assetTypes)

}

func RegisterAssetTypeRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	AssetTypeGroup := rg.Group("/asset/type", handle.MiddlewareFunc())
	{
		AssetTypeGroup.POST("/create", CreateAssetType)
		AssetTypeGroup.GET("/one/:id", GetAssetTypeById)
		AssetTypeGroup.DELETE("/void/:id", InactiveAssetType)
		AssetTypeGroup.POST("/update", UpdateAssetType)
		AssetTypeGroup.GET("/all", GetAssetTypes)
		AssetTypeGroup.POST("/list", ListPageAssetTypes)
	}
}
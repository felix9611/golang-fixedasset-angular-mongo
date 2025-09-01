package controllers

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"golang-fixedasset-mongo-backend/backend/dto"
	"net/http"
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
)


func CreateAssetList(c *gin.Context) {
	var assetItem models.AssetLists
	if err := c.ShouldBindJSON(&assetItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to create the asset item
	createdAsset, err := services.CreateAssetItem(&assetItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdAsset)
}

func ListAssetItems(c *gin.Context) {
	var req dto.ListAssetReqDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to list the asset items
	assetItems, err := services.ListAssetItems(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assetItems)
}

func GetOneAssetItem(c *gin.Context) {
	id := c.Param("id")
	assetItem, err := services.GetOneAssetItemByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, assetItem)
}

func UpdateAssetItem(c *gin.Context) {
	var updateData models.AssetLists
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to update the asset item
	updatedAsset, err := services.UpdateAssetItem(&updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAsset)
}

func RegisterAssetListRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	assetListGroup := rg.Group("/asset/asset-list", handle.MiddlewareFunc())
	{
		assetListGroup.POST("/create", CreateAssetList)
		assetListGroup.POST("/list", ListAssetItems)
		assetListGroup.GET("/one/:id", GetOneAssetItem)
		assetListGroup.POST("/update", UpdateAssetItem)

	}
}
package controllers

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	//"golang-fixedasset-mongo-backend/backend/dto"
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

func RegisterAssetListRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	assetListGroup := rg.Group("/asset/asset-list", handle.MiddlewareFunc())
	{
		assetListGroup.POST("/create", CreateAssetList)
	}
}
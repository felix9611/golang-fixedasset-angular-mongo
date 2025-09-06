package controllers

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"golang-fixedasset-mongo-backend/backend/dto"
	"net/http"
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
)

func CreateStockTakeApi(c *gin.Context) {
	var stockTake models.StockTakes
	if err := c.ShouldBindJSON(&stockTake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := services.CreateStockTakeForm(&stockTake)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, data)
}

func ListPageStockTake(c *gin.Context) {
	var query dto.ListStockTakeDto
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stockTakes, err := services.ListStockTakeForms(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list"})
		return
	}
	c.JSON(http.StatusOK, stockTakes)
}

func GetStockTakeById(c *gin.Context) {
	id := c.Param("id")
	data, err := services.GetOneStockTakeFormByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stock take form"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func StockTakeItemSubmit(c *gin.Context) {
	var stockTakeItem models.StockTakeItems
	if err := c.ShouldBindJSON(&stockTakeItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := services.StockTakeItemSubmit(&stockTakeItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit stock take item"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func RegisterStockTakeRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	stockTakeGroup := rg.Group("/asset/stock-take", handle.MiddlewareFunc())
	{
		stockTakeGroup.POST("/create", CreateStockTakeApi)
		stockTakeGroup.POST("/list", ListPageStockTake)
		stockTakeGroup.GET("/one/:id", GetStockTakeById)
		stockTakeGroup.POST("/item-submit", StockTakeItemSubmit)
	}
}
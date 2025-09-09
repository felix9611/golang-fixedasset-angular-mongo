package controllers

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"golang-fixedasset-mongo-backend/backend/dto"
	"net/http"
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
)

func CreateRepairRecord(c *gin.Context) {
	var record models.RepairRecords

	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreateRepairRecord(&record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ListPageRepairRecords(c *gin.Context) {
	var pageDto dto.RepairRecordPageReqDTO
	if err := c.ShouldBindJSON(&pageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.ListRepairRecords(&pageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetOneRepairRecord(c *gin.Context) {
	id := c.Param("id")

	result, err := services.GetOneRepairRecord(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateRepairRecord(c *gin.Context) {
	var record models.RepairRecords
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateRepairRecord(&record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func VoidRepairRecord(c *gin.Context) {
	id := c.Param("id")
	result, err := services.VoidRepairRecord(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func RegisterRepairRecordRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	repairRecordRoute := rg.Group("/asset/repair-record", handle.MiddlewareFunc())
	{
		repairRecordRoute.POST("/create", CreateRepairRecord)
		repairRecordRoute.POST("/list", ListPageRepairRecords)
		repairRecordRoute.GET("/one/:id", GetOneRepairRecord)
		repairRecordRoute.POST("/update", UpdateRepairRecord)
		repairRecordRoute.DELETE("/void/:id", VoidRepairRecord)
	}
}
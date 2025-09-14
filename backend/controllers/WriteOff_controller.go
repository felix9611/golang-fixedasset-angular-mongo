package controllers

import (
	"golang-fixedasset-mongo-backend/backend/dto"
	"golang-fixedasset-mongo-backend/backend/services"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func CreateWriteOffRecord(c *gin.Context) {
	var writeOffRecordDto dto.CreateWriteOffRecrod
	if err := c.ShouldBindJSON(&writeOffRecordDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreateWriteOff(writeOffRecordDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ListPageWriteOff(c *gin.Context) {
	var writeOffPageDto dto.ListWriteOffReqDto
	if err := c.ShouldBindJSON(&writeOffPageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.ListPageWriteOff(writeOffPageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list write off records"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func RegisterWriteOffsRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	writeOffGroup := rg.Group("/asset/write-off", handle.MiddlewareFunc())
	{
		writeOffGroup.POST("/create", CreateWriteOffRecord)
		writeOffGroup.POST("/list", ListPageWriteOff)
	}
}

package controllers

import (
	"golang-fixedasset-mongo-backend/backend/services"
	"golang-fixedasset-mongo-backend/backend/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	jwt "github.com/appleboy/gin-jwt/v2"
)

func ListInvRecord_(c *gin.Context) {
	var req dto.ListRecordReqDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.ListInvRecords(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func RegisterInvRecordRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	InvRecordGroup := rg.Group("/sys/inv-record", handle.MiddlewareFunc())
	{
		InvRecordGroup.POST("/list", ListInvRecord_)
	}
}
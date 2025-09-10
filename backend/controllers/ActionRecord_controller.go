package controllers

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
	"golang-fixedasset-mongo-backend/backend/services"
	"golang-fixedasset-mongo-backend/backend/dto"
	"net/http"
)

func ListPageActionRecords(c *gin.Context) {
	var pageDto dto.ListActionRecordReqDto
	if err := c.ShouldBindJSON(&pageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	actionRecords, err := services.ListActionRecords(&pageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list action records"})
		return
	}
	c.JSON(http.StatusOK, actionRecords)
}

func RegisterActionRecordsRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	ActionRecordGroup := r.Group("/action-records", authMiddleware.MiddlewareFunc())
	{
		ActionRecordGroup.POST("/list", ListPageActionRecords)
	}
}
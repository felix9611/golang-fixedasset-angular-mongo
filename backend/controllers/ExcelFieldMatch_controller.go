package controllers

import (
	"golang-fixedasset-mongo-backend/backend/dto"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func ListPageExcelFieldMatch(c *gin.Context) {
	var param dto.ExcelFieldMatchPageDto
	result, err := services.ExcelFieldMatchListAndPage(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateExcelFieldMatch(c *gin.Context) {
	var param models.ExcelFieldMatchs
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.CreateExcelFieldMatch(&param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetOneExcelFieldMatchById(c *gin.Context) {
	id := c.Param("id")
	result, err := services.GetOneExcelFieldMatch(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func VoidOneExcelFieldMatchById(c *gin.Context) {
	id := c.Param("id")
	result, err := services.InactiveOneExcelFieldMatch(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdateExcelFieldMatchById(c *gin.Context) {
	var param models.ExcelFieldMatchs
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.UpdateOneExcelFieldMatch(&param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func RegisterExcelFieldMatchRoutes(r *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	excelFieldMatchGroup := r.Group("/sys/excel-field-match", handle.MiddlewareFunc())
	{
		excelFieldMatchGroup.GET("/one/:id", GetOneExcelFieldMatchById)
		excelFieldMatchGroup.POST("/list", ListPageExcelFieldMatch)
		excelFieldMatchGroup.POST("/create", CreateExcelFieldMatch)
		excelFieldMatchGroup.DELETE("/void/:id", VoidOneExcelFieldMatchById)
		excelFieldMatchGroup.POST("/update", UpdateExcelFieldMatchById)
	}
}

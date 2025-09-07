package controllers

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"golang-fixedasset-mongo-backend/backend/dto"
	"net/http"
)

func CreateBudgetRecord(c *gin.Context) {
	var budgetDto models.Budgets
	if err := c.ShouldBindJSON(&budgetDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreateBudgetRecord(&budgetDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetBudgetRecordById(c *gin.Context) {
	id := c.Param("id")
	result, err := services.GetOneBudgetRecord(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func ListPageBudgetRecords(c *gin.Context) {
	var budgetPageDto dto.ListBudgetRecordsDto
	if err := c.ShouldBindJSON(&budgetPageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	budgets, err := services.ListBudgetRecords(&budgetPageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list budgets"})
		return
	}
	c.JSON(http.StatusOK, budgets)
}


func GetBudgetSummary(c *gin.Context) {
	result, err := services.GetBudgetSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdateBudgetRecord(c *gin.Context) {
	var budgetDto models.Budgets
	if err := c.ShouldBindJSON(&budgetDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateBudgetRecord(&budgetDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func InactiveBudgetRecord(c *gin.Context) {
	id := c.Param("id")
	result, err := services.VoidBudgetRecord(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Budget record inactivated successfully", "data": result})
}

func RegisterBudgetRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	BudgetGroup := rg.Group("/base/budget", handle.MiddlewareFunc())
	{
		BudgetGroup.POST("/create", CreateBudgetRecord)
		BudgetGroup.GET("/one/:id", GetBudgetRecordById)
		BudgetGroup.POST("/list", ListPageBudgetRecords)
		BudgetGroup.GET("/getBudgetSummary", GetBudgetSummary)
		BudgetGroup.POST("/update", UpdateBudgetRecord)
		BudgetGroup.DELETE("/void/:id", InactiveBudgetRecord)
	}
}

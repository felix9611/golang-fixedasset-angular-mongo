package controllers

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"golang-fixedasset-mongo-backend/backend/dto"
	"net/http"
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
)

func CreateTaxInformation(c *gin.Context) {
	var taxInfo models.TaxInformations
	if err := c.ShouldBindJSON(&taxInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreateTaxInformation(&taxInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetOneTaxInformation(c *gin.Context) {
	id := c.Param("id")
	result, err := services.GetOneTaxInformation(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func VoidOneTaxInformation(c *gin.Context) {
	id := c.Param("id")
	result, err := services.VoidOneTaxInformation(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tax Information inactivated successfully", "data": result })
}

func ListTaxInformation(c *gin.Context) {
	var taxInfoPageDto dto.TaxInformationListDTO
	if err := c.ShouldBindJSON(&taxInfoPageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.TaxInformationList(&taxInfoPageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateTaxInformationById(c *gin.Context) {
	var taxInfo models.TaxInformations
	if err := c.ShouldBindJSON(&taxInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateTaxInformation(&taxInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}


func ListTaxInformations(c *gin.Context) {
	var taxInfoPageDto dto.TaxInformationListDTO
	if err := c.ShouldBindJSON(&taxInfoPageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.TaxInformationList(&taxInfoPageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ListAllTaxInformation(c *gin.Context) {
	result, err := services.ListAllTaxInformation()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func RegisterTaxInformationRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	TaxInformationRoute := rg.Group("/base/tax-information", handle.MiddlewareFunc())
	{
		TaxInformationRoute.POST("/create", CreateTaxInformation)
		TaxInformationRoute.GET("/one/:id", GetOneTaxInformation)
		TaxInformationRoute.POST("/update", UpdateTaxInformationById)
		TaxInformationRoute.DELETE("/void/:id", VoidOneTaxInformation)
		TaxInformationRoute.POST("/list", ListTaxInformations)
		TaxInformationRoute.GET("/all", ListAllTaxInformation)
	}

}

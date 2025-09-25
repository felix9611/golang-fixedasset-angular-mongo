package controllers

import (
	"golang-fixedasset-mongo-backend/backend/dto"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func CreateLocation(c *gin.Context) {
	var location models.Locations
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreateLocation(&location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary      Get one location record by id
// @Description  Get one location record by id
// @Tags         Location
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Location ID"
// @Success      200      {object}  models.Locations
// @Router       /base/location/one/{id} [get]
func GetLocationById(c *gin.Context) {
	id := c.Param("id")
	result, err := services.GetOneLocationById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func InactiveLocationByID(c *gin.Context) {
	id := c.Param("id")
	result, err := services.InactiveLocationByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deactivate role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role deactivated successfully", "id": id, "data": result})
}

func UpdateLocationById(c *gin.Context) {
	var location models.Locations
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateLocationById(&location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ListLocations(c *gin.Context) {
	var locationPageDto dto.LocationPageDto
	if err := c.ShouldBindJSON(&locationPageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.LocationList(&locationPageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary      Get all location record
// @Description  Get all location record
// @Tags         Location
// @Produce      json
// @Success      200      {object}  []models.Locations
// @Router       /base/location/all [get]
func GetAllLocations(c *gin.Context) {
	locations, err := services.ListAllLocation()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get locations"})
		return
	}
	c.JSON(http.StatusOK, locations)
}

func RegisterLocationRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	locationGroup := rg.Group("/base/location", handle.MiddlewareFunc())
	{
		locationGroup.POST("/create", CreateLocation)
		locationGroup.GET("/one/:id", GetLocationById)
		locationGroup.POST("/update", UpdateLocationById)
		locationGroup.DELETE("/void/:id", InactiveLocationByID)
		locationGroup.POST("/list", ListLocations)
		locationGroup.GET("/all", GetAllLocations)
	}
}

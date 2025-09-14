package controllers

import (
	"golang-fixedasset-mongo-backend/backend/dto"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// @Summary      Get one vender record by id
// @Description  Get one vender record by id
// @Tags         Vendor
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Vendor ID"
// @Success      200      {object}  models.Vendors
// @Router       /base/vendor/one/{id} [get]
func GetOneVendorByIdGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendor ID"})
		return
	}

	vendor, err := services.GetOneVendorById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve vendor"})
		return
	}

	c.JSON(http.StatusOK, vendor)
}

func CreateVendorPost(c *gin.Context) {
	var vendor models.Vendors
	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreateVendor(&vendor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create vendor"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result})
}

func DeleteVendorById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendor ID"})
		return
	}

	result, err := services.InactiveVendorByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete vendor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func VendorList(c *gin.Context) {
	var pageDto dto.VendorPageDto
	if err := c.ShouldBindJSON(&pageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.VendorList(&pageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve vendor list"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateVendor(c *gin.Context) {
	var vendor models.Vendors
	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateVendorById(&vendor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vendor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func GetAllVendors(c *gin.Context) {
	vendors, err := services.GetAllVendors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get vendors"})
		return
	}
	c.JSON(http.StatusOK, vendors)
}

func RegisterVendorRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	vendorGroup := rg.Group("/base/vendor", handle.MiddlewareFunc())
	{
		vendorGroup.GET("/one/:id", GetOneVendorByIdGet)
		vendorGroup.POST("/create", CreateVendorPost)
		vendorGroup.DELETE("/void/:id", DeleteVendorById)
		vendorGroup.POST("/list", VendorList)
		vendorGroup.POST("/update", UpdateVendor)
		vendorGroup.GET("/all", GetAllVendors)
	}
}

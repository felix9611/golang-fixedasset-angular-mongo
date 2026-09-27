package controllers

import (
	"golang-fixedasset-mongo-backend/backend/dto"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"net/http"

	//	"time"
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// var tongsCollection = config.GetCollection("tongs")

var d models.Department

// @Summary      Get one Department record by id
// @Description  Get one Department record by id
// @Tags         Department
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Department ID"
// @Success      200      {object}  models.Department
// @Router       /sys/department/{id} [get]
func GetOneDepartmentById(c *gin.Context) {
	id := c.Param("id")
	department, err := services.GetOneDepartment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get department"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": department})
}

// @Summary      Get all Department record
// @Description  Get all Department record
// @Tags         Department
// @Produce      json
// @Success      200      {object}  []models.Department
// @Router       /sys/department/all [get]
func GetAllDepartments(c *gin.Context) {
	departments, err := services.GetAllDepartments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get departments"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": departments})
}

// @Summary      Create one dpartment record
// @Description  Create one dpartment record
// @Tags         Department
// @Accept       json
// @Produce      json
// @Param        request  body      models.Department  true  "Department Request Body"
// @Success      200      {object}  dto.GeneralCreateResponseBody
// @Router      /sys/dpartment/create [post]
func CreateDepartment(c *gin.Context) {
	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.CreateDepartment(&department)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// @Summary      Update one dpartment record
// @Description  Update one dpartment record
// @Tags         Department
// @Accept       json
// @Produce      json
// @Param        request  body      models.Department  true  "Department Request Body for update"
// @Success      200      {object}  dto.GeneralUpdateResponse
// @Router      /sys/dpartment/update/{id} [post]
func UpdateDepartment(c *gin.Context) {
	id := c.Param("id")
	log.Println("UpdateDepartment ID:", id)
	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateDeptById(id, &department)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update department"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Tongs by ID", "id": id, "data": result})
}

// @Summary      Delete one dpartment record by id
// @Description  Delete one dpartment record by id
// @Tags         Department
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Department ID"
// @Success      200      {object}  dto.GeneralUpdateInactiveResponseBody
// @Router      /sys/dpartment/void/{id} [delete]
func VoidDepartmentById(c *gin.Context) {
	id := c.Param("id")
	result, err := services.VoidDepartmentById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to void department"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Department voided successfully", "data": result})
}

// @Summary      List Departments
// @Description  List Departments
// @Tags         Department
// @Accept       json
// @Produce      json
// @Param        request  body      dto.DepartmentPageDto  true  "List Action Record Request Body"
// @Success      200      {object}  dto.DepartmentList
// @Router       /sys/department/list [post]
func ListPageDepartment(c *gin.Context) {
	var deptPageDto dto.DepartmentPageDto
	if err := c.ShouldBindJSON(&deptPageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	departments, err := services.DepartmentList(&deptPageDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list departments"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": departments})

}

func RegisterDepartmentRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	departments := rg.Group("/sys/department", handle.MiddlewareFunc())
	{
		departments.GET("/one/:id", GetOneDepartmentById)
		departments.POST("/create", CreateDepartment)
		departments.POST("/update/:id", UpdateDepartment)
		departments.DELETE("/void/:id", VoidDepartmentById)
		departments.POST("/list", ListPageDepartment)
		departments.GET("/all", GetAllDepartments)
	}
}

package controllers

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"golang-fixedasset-mongo-backend/backend/dto"
	"net/http"
//	"time"

	"github.com/gin-gonic/gin"
	"log"

)

// var tongsCollection = config.GetCollection("tongs")

var d models.Department

func GetOneDepartmentById(c *gin.Context) {
	id := c.Param("id")
	department, err := services.GetOneDepartment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get department"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": department })
}

func GetAllDepartments(c *gin.Context) {
	departments, err := services.GetAllDepartments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get departments"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": departments})
}

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

func VoidDepartmentById(c *gin.Context) {
	id := c.Param("id")
	result, err := services.VoidDepartmentById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to void department"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Department voided successfully", "data": result})
}

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

func RegisterDepartmentRoutes(rg *gin.RouterGroup) {
	departments := rg.Group("/sys/department")
	{
		departments.GET("/one/:id", GetOneDepartmentById)
		departments.POST("/create", CreateDepartment)
		departments.POST("/update/:id", UpdateDepartment)
		departments.DELETE("/void/:id", VoidDepartmentById)
		departments.POST("/list", ListPageDepartment)
		departments.GET("/all", GetAllDepartments)
	}
}

package controllers

import (
	"golang-fixedasset-mongo-backend/backend/auth"
	"golang-fixedasset-mongo-backend/backend/dto"
	"golang-fixedasset-mongo-backend/backend/services"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// @Summary      List Inventory Records
// @Description  List Inventory Records
// @Tags         Inventory Record
// @Accept       json
// @Produce      json
// @Param        request  body      dto.InvRecordListDto  true  "List Action Record Request Body"
// @Success      200      {object}  dto.InvRecordListResponse  "List Inventory Records"
// @Router       /sys/inv-record/list [post]
func ListInvRecord_(c *gin.Context) {
	value, _ := c.Get(identityKey)
	user, ok := value.(*auth.AuthUser)
	if !ok || user == nil {
		c.JSON(401, gin.H{"error": "Invalid user"})
		return
	}

	if !user.SuperAdmin {
		c.JSON(403, gin.H{
			"error": "Access denied",
		})
	} else {
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
}

func ListInvRecordWithFilter_(c *gin.Context) {
	value, _ := c.Get(identityKey)
	user, ok := value.(*auth.AuthUser)
	if !ok || user == nil {
		c.JSON(401, gin.H{"error": "Invalid user"})
		return
	}

	if !user.SuperAdmin {
		c.JSON(403, gin.H{
			"error": "Access denied",
		})
	} else {
		var req dto.ListRecordReqDto

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := services.ListInvRecordsWithFilter(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func RegisterInvRecordRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	InvRecordGroup := rg.Group("/sys/inv-record", handle.MiddlewareFunc())
	{
		InvRecordGroup.POST("/list", ListInvRecord_)
		InvRecordGroup.POST("/filter/list", ListInvRecordWithFilter_)
	}
}

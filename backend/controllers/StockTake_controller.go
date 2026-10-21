package controllers

import (
	"golang-fixedasset-mongo-backend/backend/auth"
	"golang-fixedasset-mongo-backend/backend/dto"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"net/http"
	"slices"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary      Create one stock take form
// @Description  Create one stock take form
// @Tags         Stock Take
// @Accept       json
// @Produce      json
// @Param        request  body      models.StockTakes  true  "List Action Record Request Body"
// @Success      200      {object}  dto.GeneralCreateResponseBody
// @Router       /asset/stock-take/create [post]
func CreateStockTakeApi(c *gin.Context) {
	var stockTake models.StockTakes
	if err := c.ShouldBindJSON(&stockTake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, _ := c.Get(identityKey)
	user, ok := value.(*auth.AuthUser)
	if !ok || user == nil {
		c.JSON(401, gin.H{"error": "Invalid user"})
		return
	}

	if !user.SuperAdmin {
		placeID, err := primitive.ObjectIDFromHex(stockTake.ActionPlaceId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !slices.Contains(user.PlaceIds, placeID) {
			c.JSON(403, gin.H{
				"error": "Location access denied",
			})
			return
		}
	}

	data, err := services.CreateStockTakeForm(&stockTake)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, data)
}

// @Summary      Update one stock take form
// @Description  Update one stock take form
// @Tags         Stock Take
// @Accept       json
// @Produce      json
// @Param        request  body      models.StockTakes  true  "List Action Record Request Body"
// @Success      200      {object}  dto.GeneralUpdateResponse
// @Router       /stock-take/update [post]
func UpdateStockTakeApi(c *gin.Context) {
	var stockTake models.StockTakes
	if err := c.ShouldBindJSON(&stockTake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, _ := c.Get(identityKey)
	user, ok := value.(*auth.AuthUser)
	if !ok || user == nil {
		c.JSON(401, gin.H{"error": "Invalid user"})
		return
	}

	if !user.SuperAdmin {
		placeID, err := primitive.ObjectIDFromHex(stockTake.ActionPlaceId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !slices.Contains(user.PlaceIds, placeID) {
			c.JSON(403, gin.H{
				"error": "Location access denied",
			})
			return
		}
	}

	data, err := services.UpdateStockTakeForm(&stockTake)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusCreated, data)
}

// @Summary      List stock take forms with pagination
// @Description  List stock take forms with pagination
// @Tags         Stock Take
// @Accept       json
// @Produce      json
// @Param        request  body      dto.ListStockTakeDto  true  "List Action Record Request Body"
// @Success      200      {object}  dto.StockTakeResponseList
// @Router       /asset/stock-take/list [post]
func ListPageStockTake(c *gin.Context) {
	var query dto.ListStockTakeDto
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stockTakes, err := services.ListStockTakeForms(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list"})
		return
	}
	c.JSON(http.StatusOK, stockTakes)
}

// @Summary      Get one stock take form by id
// @Description  Get one stock take form by id
// @Tags         Stock Take
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Stock Take Form ID"
// @Success      200      {object}  dto.GetStockTakeFormWithItemsResponse
// @Router       /asset/stock-take/one/{id} [get]
func GetStockTakeById(c *gin.Context) {
	id := c.Param("id")
	data, err := services.GetOneStockTakeFormByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stock take form"})
		return
	}

	value, _ := c.Get(identityKey)
	user, ok := value.(*auth.AuthUser)
	if !ok || user == nil {
		c.JSON(401, gin.H{"error": "Invalid user"})
		return
	}

	if !user.SuperAdmin {
		stockTake, ok := data.(map[string]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stock take form"})
			return
		}
		actionPlaceID, ok := stockTake["actionPlaceId"].(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stock take form"})
			return
		}
		placeID, err := primitive.ObjectIDFromHex(actionPlaceID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !slices.Contains(user.PlaceIds, placeID) {
			c.JSON(403, gin.H{
				"error": "Location access denied",
			})
			return
		}
	}
	c.JSON(http.StatusOK, data)
}

// @Summary      Submit stock take item
// @Description  Submit stock take item
// @Tags         Stock Take
// @Accept       json
// @Produce      json
// @Param        request  body      models.StockTakeItems  true  "Stock Take Item Request Body"
// @Success      200      {object}  dto.GeneralCreateResponseBody
// @Router       /asset/stock-take/item-submit [post]
func StockTakeItemSubmit(c *gin.Context) {
	var stockTakeItem models.StockTakeItems
	if err := c.ShouldBindJSON(&stockTakeItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := services.StockTakeItemSubmit(&stockTakeItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit stock take item"})
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary      Finish stock take form
// @Description  Finish stock take form
// @Tags         Stock Take
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Stock Take Form ID"
// @Success      200      {object}  dto.GeneralUpdateResponse
// @Router       /asset/stock-take/finish/{id} [delete]
func FinishStockTake(c *gin.Context) {
	id := c.Param("id")
	data, err := services.FinishOrVoidStockTakeForm(id, 2, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to finish stock take form"})
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary      Void stock take form
// @Description  Void stock take form
// @Tags         Stock Take
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Stock Take Form ID"
// @Success      200      {object}  dto.GeneralCreateResponseBody
// @Router       /asset/stock-take/void/{id} [delete]
func VoidStockTake(c *gin.Context) {
	id := c.Param("id")
	data, err := services.FinishOrVoidStockTakeForm(id, 0, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to finish stock take form"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func RegisterStockTakeRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	stockTakeGroup := rg.Group("/asset/stock-take", handle.MiddlewareFunc())
	{
		stockTakeGroup.POST("/create", CreateStockTakeApi)
		stockTakeGroup.POST("/list", ListPageStockTake)
		stockTakeGroup.GET("/one/:id", GetStockTakeById)
		stockTakeGroup.POST("/item-submit", StockTakeItemSubmit)
		stockTakeGroup.DELETE("/finish/:id", FinishStockTake)
		stockTakeGroup.DELETE("/void/:id", VoidStockTake)
		stockTakeGroup.POST("/update", UpdateStockTakeApi)
	}
}

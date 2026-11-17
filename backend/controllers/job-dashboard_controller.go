package controllers

import (
	"golang-fixedasset-mongo-backend/backend/auth"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/services"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// @Summary      Create job task
// @Description  Create job task
// @Tags         Task Board
// @Accept       json
// @Produce      json
// @Param        request  body      models.JobDashboard  true  "Job Dashboard Request Body"
// @Success      200      {object}  dto.GeneralCreateResponseBody
// @Router      /tasks/create [post]
func CreateJobBoard(c *gin.Context) {
	value, _ := c.Get(identityKey)
	user, ok := value.(*auth.AuthUser)
	if !ok || user == nil {
		c.JSON(401, gin.H{"error": "Invalid user"})
		return
	}

	var task models.JobDashboard
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !(user.SuperAdmin || user.Manager) && task.Status >= 3 {
		c.JSON(403, gin.H{
			"error": "You do not allow to create this task on this status",
		})
	} else {
		id, err := services.CreateTasks(&task)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

func RegisterTaskBoardsRoutes(rg *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	taskBoards := rg.Group("/tasks", handle.MiddlewareFunc())
	{
		taskBoards.POST("/create", CreateJobBoard)
	}
}

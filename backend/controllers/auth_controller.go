package controllers

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
	"golang-fixedasset-mongo-backend/backend/auth"
	"golang-fixedasset-mongo-backend/backend/services"
)

var (
	identityKey = "username" 
	port        string
)


func HeyTongsHandler(c *gin.Context) {
    userRaw, exists := c.Get(identityKey)
    if !exists {
        c.JSON(401, gin.H{
            "message": "Unauthorized: identity not found",
        })
        return
    }

    user, ok := userRaw.(*auth.AuthUser)
    if !ok {
        c.JSON(401, gin.H{
            "message": "Unauthorized: invalid identity type",
        })
        return
    }

    c.JSON(200, gin.H{
        "username": user.Username,
        "text":     "Hey Tongs.",
    })
}

func GetUserDetails(c *gin.Context) {
	userRaw, exists := c.Get(identityKey)
	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized: identity not found",
		})
		return
	}

	authUser, ok := userRaw.(*auth.AuthUser)
	if !ok {
		c.JSON(401, gin.H{
			"message": "Unauthorized: invalid identity type",
		})
		return
	}

	username := authUser.Username
	userDetails, err := services.GetUserDetailByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error: failed to get user details",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": userDetails,
	})
}


func RegisterAuthRoute(r *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	r.POST("/login", handle.LoginHandler)
	auth := r.Group("/auth", handle.MiddlewareFunc())
	{
		auth.GET("/refresh_token", handle.RefreshHandler)
		auth.GET("/verify-token", HeyTongsHandler)
		auth.GET("/user-profile", GetUserDetails)
	}
}
package controllers

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
	"golang-fixedasset-mongo-backend/backend/auth"
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

func RegisterAuthRoute(r *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	r.POST("/login", handle.LoginHandler)
	auth := r.Group("/auth", handle.MiddlewareFunc())
	{
		auth.GET("/refresh_token", handle.RefreshHandler)
		auth.GET("/test", HeyTongsHandler)
	}
}
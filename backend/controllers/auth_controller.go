package controllers

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"
)

func RegisterAuthRoute(r *gin.RouterGroup, handle *jwt.GinJWTMiddleware) {
	r.POST("/login", handle.LoginHandler)
}
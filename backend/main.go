package main

import (
	"github.com/gin-contrib/cors"
	"golang-fixedasset-mongo-backend/backend/config"
	"golang-fixedasset-mongo-backend/backend/controllers"
	"golang-fixedasset-mongo-backend/backend/example"
	"golang-fixedasset-mongo-backend/backend/auth"
	"github.com/gin-gonic/gin"
	"time"
	jwt "github.com/appleboy/gin-jwt/v2"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))


	config.ConnectDatabase()

	authMiddleware, errAuth := jwt.New(example.InitParams())
	if errAuth != nil {
		panic("JWT Middleware initialization failed: " + errAuth.Error())
	}

	authServiceMiddleware, errAuthService := jwt.New(auth.InitAuthParams())
	if  errAuthService != nil {
		panic("JWT Middleware initialization failed: " + errAuth.Error())
	}

	api := router.Group("/")
	controllers.RegisterTongsRoutes(api)
	controllers.RegisterDepartmentRoutes(api)
	controllers.RegisterSysUserRoutes(api)
	example.RegisterAuthExampleRoute(api, authMiddleware)
	controllers.RegisterAuthRoute(api, authServiceMiddleware)
	controllers.RegisterSysRoleRoutes(api)

	router.Run(":6500")
}

package main

import (
	"golang-fixedasset-mongo-backend/backend/config"
	"golang-fixedasset-mongo-backend/backend/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDatabase()

	api := router.Group("/api")
	controllers.RegisterTongsRoutes(api)

	router.Run(":6500")
}

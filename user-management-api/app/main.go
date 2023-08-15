package main

import (
	"app/configs"
	"app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Run database
	configs.ConnectDB()

	// Routes
	routes.UserRoute(router)

	router.Run("localhost:6000")
}

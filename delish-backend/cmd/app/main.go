package main

import (
	"delish-backend/internal/initializers"
	dish_routes "delish-backend/internal/routes/v1"

	"github.com/gin-gonic/gin"
)

func init() {
	// Load environment variables
	initializers.LoadEnvs()

	// Initialize the database
	initializers.InitDB()
}

func main() {
	// Initialize the application
	app := gin.Default()

	// Use the routes
	dish_routes.Use(app)

	// Run the application
	app.Run()
}

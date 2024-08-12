package main

import (
	"delish-backend/internal/config"
	"delish-backend/internal/initializers"
	dishes_routes "delish-backend/internal/routes/v1"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
    // Load environment variables
    initializers.LoadEnvs()

    // Initialize the database
    initializers.InitDB()
}

func main() {
    // Set Gin mode based on GO_ENV value
    if os.Getenv("GO_ENV") == "production" {
        gin.SetMode(gin.ReleaseMode)
    } else {
        gin.SetMode(gin.DebugMode)
    }

    // Initialize the application
    app := gin.New()

    // Set trusted proxies (only for production)
    if os.Getenv("GO_ENV") == "production" {
        // Set list of trusted proxies
        err := app.SetTrustedProxies([]string{
            "192.168.1.1",
            "192.168.1.2",
        })

        // Panic if any error occurred
        if err != nil {
            panic(err)
        }
    }

    // Use the default recovery middleware for error handling
    app.Use(gin.Recovery())

    // Configure CORS middleware
    app.Use(cors.New(config.CORS))

    // Use the routes
    dishes_routes.RegisterRoutes(app)

    // Run the application
    app.Run()
}
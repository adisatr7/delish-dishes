package v1

import (
	controllers "delish-backend/internal/controllers/v1"

	"github.com/gin-gonic/gin"
)

func Order(app *gin.Engine) {
    // Group all routes under `/api/v1/orders`
    route := app.Group("/api/v1/orders")
    {
        // Route to create a new order
        route.POST("/", controllers.CreateDish)
        // Route to get all orders
        route.GET("/", controllers.GetAllOrders)

        // Group routes that operate on a specific dish identified by its ID
        specificRoute := route.Group("/:id")
        {
            // Route to get a specific dish by ID
            specificRoute.GET("/", controllers.GetOrderById)
            // Route to update a specific dish by ID
            specificRoute.PATCH("/", controllers.UpdateOrderById)
            // Route to delete a specific dish by ID
            specificRoute.DELETE("/", controllers.DeleteOrderById)
        }
    }
}

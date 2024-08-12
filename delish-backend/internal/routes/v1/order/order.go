package order

import (
	order_controllers "delish-backend/internal/controllers/v1/order"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
    // Group all routes under `/api/v1/orders`
    route := app.Group("/api/v1/orders")
    {
        // Route to create a new order
        route.POST("/", order_controllers.CreateOrder)
        // Route to get all orders
        route.GET("/", order_controllers.GetAllOrders)

        // Group routes that operate on a specific order identified by its ID
        specificRoute := route.Group("/:id")
        {
            // Route to get a specific order by ID
            specificRoute.GET("/", order_controllers.GetOrderById)
            // Route to update a specific order by ID
            specificRoute.PATCH("/", order_controllers.UpdateOrderById)
            // Route to delete a specific order by ID
            specificRoute.DELETE("/", order_controllers.DeleteOrderById)
        }
    }
}

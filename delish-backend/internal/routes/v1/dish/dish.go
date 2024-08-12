package dish

import (
	dish_controllers "delish-backend/internal/controllers/v1/dish"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
    // Group all routes under `/api/v1/dishes``
    route := app.Group("/api/v1/dishes")
    {
        // Route to create a new dish
        route.POST("/", dish_controllers.CreateDish)
        // Route to get all dishes
        route.GET("/", dish_controllers.GetAllDishes)

        // Group routes that operate on a specific dish identified by its ID
        specificRoute := route.Group("/:id")
        {
            // Route to get a specific dish by ID
            specificRoute.GET("/", dish_controllers.GetDishById)
            // Route to update a specific dish by ID
            specificRoute.PATCH("/", dish_controllers.UpdateDishById)
            // Route to delete a specific dish by ID
            specificRoute.DELETE("/", dish_controllers.DeleteDishById)
        }
    }
}
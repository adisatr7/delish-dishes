package v1

import (
	controllers "delish-backend/internal/controllers/v1"

	"github.com/gin-gonic/gin"
)

func Dish(app *gin.Engine) {
    // Group all routes under `/api/v1/dishes``
    route := app.Group("/api/v1/dishes")
    {
        // Route to create a new dish
        route.POST("/", controllers.CreateDish)
        // Route to get all dishes
        route.GET("/", controllers.GetAllDishes)

        // Group routes that operate on a specific dish identified by its ID
        specificRoute := route.Group("/:id")
        {
            // Route to get a specific dish by ID
            specificRoute.GET("/", controllers.GetDishById)
            // Route to update a specific dish by ID
            specificRoute.PATCH("/", controllers.UpdateDishById)
            // Route to delete a specific dish by ID
            specificRoute.DELETE("/", controllers.DeleteDishById)
        }
    }
}
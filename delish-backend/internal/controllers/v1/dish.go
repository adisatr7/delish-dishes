package v1

import (
	"delish-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a new Dish record
func CreateDish(ctx *gin.Context) {
	// Parse the request body
	var body struct {
		Name string
		Desc string
	}
	ctx.BindJSON(&body)

	// Validate the request body
	if body.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Name is required",
		})
		return
	}

	// Create a new Dish record
	err := services.CreateDish(services.Params{
		Name: body.Name,
		Desc: body.Desc,
	})

	// On error, return a 500 Internal Server Error response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// On success, return a 201 Created response
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Dish created successfully",
	})
}

// Get all Dish records
func GetAllDish(ctx *gin.Context) {
	// Get all Dish records
	dishes, err := services.GetAllDish()

	// On error, return a 500 Internal Server Error response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// On success, return a 200 OK response with the Dish records
	ctx.JSON(http.StatusOK, dishes)
}

// Get one Dish record by its ID
func GetDishById(ctx *gin.Context) {
	// Get the Dish ID from the URL
	dish
}

// Update a Dish record by its ID
func PatchDishById(ctx *gin.Context) {
	// body := ctx.Request.Body
}

// Delete a Dish record by its ID
func DeleteDishById(ctx *gin.Context) {
	// body := ctx.Request.Body
}

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
		Name  string
		Desc  string
		Price *int64
	}
	ctx.BindJSON(&body)

	// Validate the request body
	if body.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Name is required!",
			"success": false,
		})
		return
	}
	if body.Price == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Price is required! If item is free, set price to 0.",
			"success": false,
		})
		return
	}
	if *body.Price < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Price is invalid!",
			"success": false,
		})
		return
	}

	// Create a new Dish record
	err := services.CreateDish(services.Params{
		Name:  body.Name,
		Desc:  body.Desc,
		Price: *body.Price,
	})

	// On error, return a 500 Internal Server Error response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// On success, return a 201 Created response
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Dish created successfully!",
		"success": true,
	})
}

// Get all Dish records
func GetAllDishes(ctx *gin.Context) {
	// Get all Dish records
	dishes, err := services.GetAllDishes()

	// On error, return a 500 Internal Server Error response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// On success, return a 200 OK response with the Dish records
	ctx.JSON(http.StatusOK, gin.H{
		"data":    dishes,
		"success": true,
	})
}

// Get one Dish record by its ID
func GetDishById(ctx *gin.Context) {
	// Get the Dish ID from the URL
	dishId := ctx.Param("id")

	// Get the Dish record by its ID
	dish, err := services.GetDishById(dishId)

	// On error, return a 500 Internal Server Error response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// On success, return a 200 OK response with the Dish record
	ctx.JSON(http.StatusOK, gin.H{
		"data":    dish,
		"success": true,
	})
}

// Update a Dish record by its ID
func UpdateDishById(ctx *gin.Context) {
	// Get the Dish ID from the URL
	dishId := ctx.Param("id")

	// Parse the request body
	var body struct {
		Name  string
		Desc  string
		Price int64
	}
	ctx.BindJSON(&body)

	// Validate the request body
	if body.Name == "" {
		// On error, return a 400 Bad Request response
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Name is required!",
			"success": false,
		})
		return
	}
	if body.Price < 0 {
		// On error, return a 400 Bad Request response
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Price is invalid!",
			"success": false,
		})
		return
	}

	// Update the Dish record by its ID
	err := services.UpdateDish(dishId, services.Params{
		Name: body.Name,
		Desc: body.Desc,
	})

	// On error, return a 500 Internal Server Error response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// On success, return a 200 OK response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Dish updated successfully!",
		"success": true,
	})
}

// Delete a Dish record by its ID
func DeleteDishById(ctx *gin.Context) {
	// Get the Dish ID from the URL
	dishId := ctx.Param("id")

	// Delete the Dish record by its ID
	err := services.DeleteDishById(dishId)

	// On error, return a 500 Internal Server Error response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// On success, return a 200 OK response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Dish deleted successfully!",
		"success": true,
	})
}

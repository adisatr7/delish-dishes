package v1

import (
	"github.com/gin-gonic/gin"
)

// Create a new Dish record
func CreateDish(ctx *gin.Context) {
	// body := ctx.Request.Body
}

// Get all Dish records
func GetAllDish(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get all dishes",
	})
}

// Get one Dish record by its ID
func GetDishById(ctx *gin.Context) {
	// body := ctx.Request.Body
}

// Update a Dish record by its ID
func PatchDishById(ctx *gin.Context) {
	// body := ctx.Request.Body
}

// Delete a Dish record by its ID
func DeleteDishById(ctx *gin.Context) {
	// body := ctx.Request.Body
}
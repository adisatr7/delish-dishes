package v1

import (
	"delish-backend/internal/initializers"

	"github.com/gin-gonic/gin"
)

var db = *initializers.DB


// Create a new Dish record
func CreateDish(ctx *gin.Context) {
	// body := ctx.Request.Body
	// db.Create(&body)
}

// Get all Dish records
func GetAllDish(ctx *gin.Context) {
	db.Find(&Dishes{})
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
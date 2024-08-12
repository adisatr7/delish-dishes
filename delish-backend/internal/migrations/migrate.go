package main

import (
	"delish-backend/internal/initializers"
	"delish-backend/internal/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.InitDB()

	if initializers.DB == nil {
		panic("Failed to connect to the database")
	}
}

func main() {
	initializers.DB.AutoMigrate(&models.DishOrder{})
	initializers.DB.AutoMigrate(&models.Dish{})
	initializers.DB.AutoMigrate(&models.Order{})
}

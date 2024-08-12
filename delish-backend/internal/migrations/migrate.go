package main

import (
    "delish-backend/internal/initializers"
    "delish-backend/internal/models"
    "log"
)

func init() {
    initializers.LoadEnvs()
    initializers.InitDB()

    // Check if the database connection is successful
    if initializers.DB == nil {
        panic("Failed to connect to the database")
    }
}

func main() {
    // Initialize tables in the correct order
    initializers.DB.AutoMigrate(&models.DishOrder{})
    initializers.DB.AutoMigrate(&models.Dish{})
    initializers.DB.AutoMigrate(&models.Order{})

    // Migration completed
    log.Println("Migration script finished running!")
}

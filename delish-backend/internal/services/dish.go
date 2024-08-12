package services

import (
	"delish-backend/internal/initializers"
	"delish-backend/internal/models"
)

// Get the database instance
var db = initializers.DB

// Define the Params struct for CREATE and UPDATE operations
type Params struct {
    Name string
    Desc string
}

// Create a new Dish record
func CreateDish(params Params) error {
	// Create a new Dish record
    dish := models.Dish{
        Name: params.Name,
        Desc: params.Desc,
    }

	// Save the new Dish record
    if err := db.Create(&dish).Error; err != nil {
		// Return the error if any
        return err
    }

	// Return nil if no error occurred
    return nil
}

// Get all Dish records
func GetAllDish() ([]*models.Dish, error) {
    // Prepare a slice to store the Dish records
    var dishes []*models.Dish

    // Get all Dish records
    if err := db.Find(&dishes).Error; err != nil {
        // Return the error if any
        return nil, err
    }

    // Return the Dish records if no error occurred
    return dishes, nil
}

// Get one Dish record by its ID
func GetDishById(id string) (*models.Dish, error) {
    // Prepare a variable to store the Dish record
    dish := &models.Dish{}

    // Get the Dish record by its ID
    if err := db.Where("id = ?", id).First(dish).Error; err != nil {
        // Return the error if any
        return nil, err
    }

    // Return the Dish record if no error occurred
    return dish, nil
}

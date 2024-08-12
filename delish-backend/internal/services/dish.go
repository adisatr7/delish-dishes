package services

import (
	"delish-backend/internal/initializers"
	"delish-backend/internal/models"
	"delish-backend/internal/utils"
	"errors"

	"gorm.io/gorm"
)

// Define the Params struct for CREATE and UPDATE operations
type Params struct {
	Name string
	Desc string
	Price int64
}

// Get the database instance
func getDbInstance() (*gorm.DB, error) {
	// Get the database instance from the initializers
	db := initializers.DB

	// Check if db is initialized
	if db == nil {
		return nil, errors.New("database is not initialized")
	}

	// Return the database instance if no error occurred
	return db, nil
}

// Create a new Dish record
func CreateDish(params Params) error {
	// Get the database instance
	db, err := getDbInstance()
	if err != nil {
		return err
	}

	// Add dot at the end of the description if it is not empty and it hasn't already
	if params.Desc != "" && params.Desc[len(params.Desc)-1] != '.' {
		params.Desc += "."
	}

	// Create a new Dish record
	dish := models.Dish{
		ID:   utils.GenerateUUID(),
		Name: params.Name,
		Desc: &params.Desc,
		Price: params.Price,
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
func GetAllDishes() ([]*models.Dish, error) {
	// Get the database instance
	db, err := getDbInstance()
	if err != nil {
		return nil, err
	}

	// Prepare a slice to store the Dish records
	var dishes []*models.Dish

	// Get all Dish records where deleted_at is NULL
	if err := db.Find(&dishes).Error; err != nil {
		// Return the error if any
		return nil, err
	}

	// Return the Dish records if no error occurred
	return dishes, nil
}

// Get one Dish record by its ID
func GetDishById(id string) (*models.Dish, error) {
	db, err := getDbInstance()
	if err != nil {
		return nil, err
	}

	// Prepare a variable to store the Dish record
	dish := &models.Dish{}

	// Get the Dish record by its ID where deleted_at is NULL
	if err := db.Where("id = ? AND deleted_at IS NULL", id).First(dish).Error; err != nil {
		// Return the error if any
		return nil, err
	}

	// Return the Dish record if no error occurred
	return dish, nil
}

// Update a Dish record by its ID
func UpdateDish(id string, params Params) error {
	db, err := getDbInstance()
	if err != nil {
		return err
	}

	// Get the Dish record
	dish, err := GetDishById(id)
	if err != nil {
		return err
	}

	// Update the Dish record with the new values
	dish.Name = params.Name
	dish.Desc = &params.Desc
	dish.Price = params.Price

	// Add dot at the end of the description if it is not empty and it hasn't already
	if dish.Desc != nil && (*dish.Desc)[len(*dish.Desc)-1] != '.' {
		*dish.Desc += "."
	}

	// Save the updated Dish record
	if err := db.Save(dish).Error; err != nil {
		return err
	}

	// Return nil if no error occurred
	return nil
}

// Soft delete a Dish record by its ID
func DeleteDishById(id string) error {
	db, err := getDbInstance()
	if err != nil {
		return err
	}

	// Get the Dish record
	dish, err := GetDishById(id)
	if err != nil {
		return err
	}

	// Soft delete the Dish record
	if err := db.Delete(dish).Error; err != nil {
		return err
	}

	// Return nil if no error occurred
	return nil
}

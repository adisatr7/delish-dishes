package utils

import (
	"delish-backend/internal/initializers"
	"errors"

	"gorm.io/gorm"
)

// Get the database instance
func GetDbInstance() (*gorm.DB, error) {
	// Get the database instance from the initializers
	db := initializers.DB

	// Check if db is initialized
	if db == nil {
		return nil, errors.New("database is not initialized")
	}

	// Return the database instance if no error occurred
	return db, nil
}
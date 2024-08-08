package initializers

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// Get the database URL from the environment variable
	dsn := os.Getenv("DB_URL")

	// Open a connection to the database
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Return the database connection to be used by the application
	return DB
}
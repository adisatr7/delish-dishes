package initializers

import (
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    // Get the database URL from the environment variable
    dsn := os.Getenv("DB_URL")
    if dsn == "" {
        log.Fatal("DB_URL environment variable is not set")
    }

    // Open a connection to the database
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
}
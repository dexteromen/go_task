package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is a global variable to hold the database connection
var DB *gorm.DB

// ConnectDatabase initializes the database connection using GORM and SQLite
func ConnectDatabase() error {
	// Open a connection to the SQLite database file "tasks.db"
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		// Log an error message if the connection fails
		log.Println("Error connecting to database:", err)
		return err
	}

	// Assign the database connection to the global variable DB
	DB = db
	return nil
}

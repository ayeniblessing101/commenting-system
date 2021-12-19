package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//NewDatabase - returns a pointer to a database object
func NewDatabase() (*gorm.DB, error) {

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return db, err
	}
	
	postgresDB, err := db.DB()

	defer postgresDB.Close()
	
	if err != nil {
		return db, err
	}

	if err := postgresDB.Ping(); err != nil {
		return db, err
	}

	fmt.Println("Setting up new Database connection")
	return db, nil
}
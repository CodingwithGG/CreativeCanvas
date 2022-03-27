package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

var db *gorm.DB
var err error

func GetDb() *gorm.DB {
	return db
}

func ConnectDatabase() error {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbpassword := os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, dbpassword, dbPort)

	// Opening connection to database
	db, err = gorm.Open(dialect, dbURI)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database successfully")
	}
	return nil
	// Close the database connection when the ConnectDatabase function closes
}

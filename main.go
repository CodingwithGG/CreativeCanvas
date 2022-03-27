package main

import (
	"awesomeProject/config"
	"awesomeProject/routes"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}
	err := config.ConnectDatabase()
	if err != nil {
		log.Println("Error Connecting to the Database", err.Error())
		return
	}
	db := config.GetDb()

	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Println("error closing database")
		} else {
			log.Println("closed database")
		}
	}(db)
	config.AutoMigrate()

	router := routes.NewRoutes()
	err = router.Run("localhost:8081")
	if err != nil {
		log.Println("Error connecting to Server", err.Error())
		return
	}
}

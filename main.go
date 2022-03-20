package main

import (
	"awesomeProject/accounts/models"
	"awesomeProject/config"
	"awesomeProject/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func main() {
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

	db.AutoMigrate(&models.User{}, models.Password{})

	router := routes.NewRoutes()
	err = router.Run("localhost:8081")
	if err != nil {
		log.Println("Error connecting to Server", err.Error())
		return
	}
}

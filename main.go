package main

import (
	"awesomeProject/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int
	PersonID   int
}

func main() {
	db := config.ConnectDatabase()
	fmt.Println("closed database")
	defer db.Close()
	//db.AutoMigrate(&Person{})
	//db.AutoMigrate(&Book{})

	//router := gin.Default()
	//router.GET("/albums", models.GetAlbums)
	//
	//router.Run("localhost:8081")

}

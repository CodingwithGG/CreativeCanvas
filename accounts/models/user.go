package models

import (
	"debug/dwarf"
	"github.com/gin-gonic/gin"
	"github.com/golang-sql/civil"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

//
//type album struct {
//	ID     string  `json:"id"`
//	Title  string  `json:"title"`
//	Artist string  `json:"artist"`
//	Price  float64 `json:"price"`
//}
//
//var albums = []album{
//	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
//}
//
//func GetAlbums(c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, albums)
//}

type Person struct {
	gorm.Model

	Name               string         `json:"name"`
	Email              string         `gorm:"typevarchar(100);unique_index"`
	VerificationStatus dwarf.EnumType `json:"verification_status"`
	OnboardingStatus   dwarf.EnumType `json:"onboarding_status"`
	IsActive           bool           `json:"is_active"`
	FirstName          string         `json:"first_name"`
	LastName           string         `json:"last_name"`
	IsStaff            bool           `json:"is_staff"`
	Gender             dwarf.EnumType `json:"gender"`
	Mobile             string         `json:"mobile"`
	ActiveAt           civil.DateTime `json:"active_at"`
	Location           time.Location  `json:"location"`
	ProfilePic         string         `json:"profile_pic"`
	RatingCount        int            `json:"rating_count"`
	RatedBy            string         `json:"rated_by"`
}

func GetPersons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persons)
}

//type Book struct {
//	gorm.Model
//
//	Title      string
//	Author     string
//	CallNumber int
//	PersonID   int
//}

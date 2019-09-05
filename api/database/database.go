package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	connectionParams := "user=postgres password=postgres sslmode=disable host=db"
	for i := 0; i < 1; i++ {
		DB, err = gorm.Open("postgres", connectionParams)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Fatalf("could not connect to database with params: %v, %v", connectionParams, err)
	}

	log.Printf("Connected to database\n")


}

func InitDB() {

	dbName := "rocket"
	DB.Exec("CREATE DATABASE " + dbName)

	if !DB.HasTable(&Booking{}) {
		DB.CreateTable(&Booking{})
	}

}
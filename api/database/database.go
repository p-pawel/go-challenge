package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	pgUser := getEnvVar("PG_USER")
	pgPass := getEnvVar("PG_PASSWORD")
	pgHost := getEnvVar("PG_HOST")
	connectionParams := "user=" + pgUser + " password=" + pgPass + " sslmode=disable host=" + pgHost
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

func getEnvVar(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("Environmental variable not set: %s", name)
	}
	return value
}

func CreateSchema() {

	dbName := "rocket"
	DB.Exec("CREATE DATABASE " + dbName)

	if !DB.HasTable(&Booking{}) {
		DB.CreateTable(&Booking{})
	}

	if !DB.HasTable(&Destination{}) {
		DB.CreateTable(&Destination{})

		createDestination("Mars")
		createDestination("Moon")
		createDestination("Pluto")
		createDestination("Asteroid Belt")
		createDestination("Europa")
		createDestination("Titan")
		createDestination("Ganymede")
	}

}

func createDestination(name string) {
	var newDestination Destination
	newDestination.Name = name
	DB.Create(&newDestination)
	//return nil
}

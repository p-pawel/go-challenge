package service

import (
	"github.com/p-pawel/go-challenge/database"
	"log"
)

func FindAllBookings() interface{} {
	var bookings []database.Booking
	database.DB.Find(&bookings)
	return bookings
}

func TryToCreateBooking(newBooking *database.Booking) []string {

	var validator DestinationValidator
	result := validator.isValid(*newBooking)
	if result != "" {
		log.Printf(result)
		return []string{result}
	}
	database.DB.Create(&newBooking)
	return nil
}



package service

import (
	"github.com/p-pawel/go-challenge/database"
	"log"
)

// TODO
// join commented out as not really useful, see the comment in entities.go
func FindAllBookings() interface{} {
	var bookings []database.Booking
	database.DB.
		//Joins("JOIN launchpads l ON l.id = launchpad_id").
		//Joins("JOIN destinations d ON d.id = destination_id").
		Find(&bookings)
	return bookings
}

var validators = []BookingValidator{
	&destinationValidator{},
	&launchpadAvailabilityValidator{},
	&overlapWithSpaceXValidator{},
}


func TryToCreateBooking(newBooking *database.Booking) []string {

	for _, v := range validators {
		result := v.isValid(*newBooking)
		if result != "" {
			log.Printf(result)
			return []string{result}
		}

	}
	database.DB.Create(&newBooking)
	return nil
}

func DeleteBooking(id uint) bool {
	db := database.DB.Delete(&database.Booking{}, "id = ?", id)
	return db.RowsAffected > 0
}

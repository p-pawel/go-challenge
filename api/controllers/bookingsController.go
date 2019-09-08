package controllers

import (
	"encoding/json"
	"github.com/p-pawel/go-challenge/database"
	"github.com/p-pawel/go-challenge/service"
	"io/ioutil"
	"net/http"
)

func GetBookings(w http.ResponseWriter, r *http.Request) {

	bookings := service.FindAllBookings()

	w.Header().Add("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(bookings)
}


func PostBooking(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var newBooking database.Booking
	{
		err := json.Unmarshal(reqBody, &newBooking)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			return
		}
	}

	{
		errors := service.TryToCreateBooking(&newBooking)
		if len(errors)>0 {
			w.WriteHeader(http.StatusUnprocessableEntity)

			w.Header().Add("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(errors)

			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newBooking)
}
package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/p-pawel/go-challenge/database"
	"github.com/p-pawel/go-challenge/service"
	"io/ioutil"
	"net/http"
	"strconv"
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

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseErr := strconv.ParseUint(vars["id"], 10, 64)

	if parseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := service.DeleteBooking(uint(id))
	if result == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	return
}

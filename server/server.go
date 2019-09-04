package server

import (
	"github.com/gorilla/mux"
	"github.com/p-pawel/go-challenge/controllers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/booking", controllers.GetBookings).Methods("GET", "POST")
	return router
}


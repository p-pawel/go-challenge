package main

import (
	"github.com/p-pawel/go-challenge/database"
	"github.com/p-pawel/go-challenge/server"
	"log"
	"net/http"
	"os"
)


func main() {

	database.ConnectDB()
	database.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := server.SetupRouter()

	if err := http.ListenAndServe(":"+port, server.LogRequest(router)); err != nil {
		log.Fatalf("could not listen on port 3000 %v", err)
	} else {
		log.Println("API server is up")
	}
}



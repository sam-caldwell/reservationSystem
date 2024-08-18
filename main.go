package main

import (
	"github.com/sam-caldwell/reservationSystem/api"
	"log"
	"net/http"
)

func main() {

	log.Println("service starting...")

	http.HandleFunc("/api/v1/reservations", api.Router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

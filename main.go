package main

import (
	"github.com/sam-caldwell/reservationSystem/api"
	"github.com/sam-caldwell/reservationSystem/ui"
	"log"
	"net/http"
)

func main() {

	log.Println("service starting...")

	http.HandleFunc("/api/v1/reservations", api.Router)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.HandleCORS(w, r)
		log.Println("index page load")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(ui.IndexHtml))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

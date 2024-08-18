package api

import (
	"github.com/sam-caldwell/reservationSystem/data"
	"log"
	"net/http"
	"time"
)

// apiReservationFetch - View existing reservation ?date=2024-08-21
func apiReservationFetch(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting apiReservationFetch()")

	dateStr := r.URL.Query().Get("date")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Printf("Error parsing date: %v", err)
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	log.Printf("Fetch reservations for %s", dateStr)

	w.Header().Set("Content-Type", "application/json")
	if err, response := data.Db.FetchReservations(date); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("{}"))
		log.Printf("Error adding reservation: %v", err)
	} else {
		if err, jsonData := response.Marshal(); err != nil {
			log.Printf("Error adding reservation: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("An error occurred"))
		} else {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(jsonData))
		}
	}
}

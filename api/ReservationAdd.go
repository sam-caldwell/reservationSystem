package api

import (
	"github.com/sam-caldwell/reservationSystem/data"
	"log"
	"net/http"
)

// apiReservationAdd - Add new reservation (from json body)
func apiReservationAdd(w http.ResponseWriter, r *http.Request) {

	log.Println("Starting apiReservationAdd()")

	var reservationRequest data.Reservation
	if err := reservationRequest.Unmarshall(r); err != nil {
		log.Printf("Error decoding json: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err, response := data.Db.AddReservation(&reservationRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("{}"))
	} else {
		if err, jsonData := response.Marshal(); err != nil {
			log.Printf("Error adding reservation: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("{}"))
		} else {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(jsonData)
		}
	}
}

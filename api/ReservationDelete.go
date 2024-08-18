package api

import (
	"github.com/sam-caldwell/reservationSystem/data"
	"log"
	"net/http"
)

// apiReservationDelete - delete an existing reservation
func apiReservationDelete(w http.ResponseWriter, r *http.Request) {

	log.Println("Starting apiReservationDelete()")

	var reservationRequest data.Reservation
	if err := reservationRequest.Unmarshall(r); err != nil {
		log.Printf("Error decoding json: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err, msg := data.Db.DeleteReservation(&reservationRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("{}"))
		log.Printf("Error adding reservation: %v", err)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(msg))
	}

}

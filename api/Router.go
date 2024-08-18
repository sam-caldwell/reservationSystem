package api

import (
	"log"
	"net/http"
)

// Router - Route API calls
func Router(w http.ResponseWriter, r *http.Request) {
	log.Printf("request routing: %v", r.Method)
	HandleCORS(w, r)
	switch r.Method {
	case http.MethodGet:
		apiReservationFetch(w, r)
	case http.MethodPut:
		apiReservationAdd(w, r)
	case http.MethodDelete:
		apiReservationDelete(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	default:
		log.Printf("unknown method: %v", r.Method)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

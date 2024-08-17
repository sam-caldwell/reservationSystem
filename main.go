package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	defaultCapacity = 4
	openHour        = 10
	closeHour       = 22
)

// Reservation - a single reservation entity
type Reservation struct {
	Name string    `json:"name"`
	Time time.Time `json:"time"` //Note: this is going to expect  RFC 3339 time (in zulu)
}

// db - a fake in-memory database for our reservations
var db = struct {
	sync.Mutex
	reservationCapacity int
	reservations        map[string][]Reservation
}{
	reservationCapacity: defaultCapacity,
	reservations:        make(map[string][]Reservation),
}

func main() {

	log.Println("service starting...")

	http.HandleFunc("/api/v1/reservations", handleRouting)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRouting(w http.ResponseWriter, r *http.Request) {
	log.Printf("request routing: %v", r.Method)
	switch r.Method {
	case http.MethodGet:
		viewReservations(w, r)
	case http.MethodPut:
		addReservation(w, r)
	case http.MethodDelete:
		cancelReservation(w, r)
	default:
		log.Printf("unknown method: %v", r.Method)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// addReservation - Add new reservation (from json body)
func addReservation(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting addReservation()")

	var res Reservation
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hour := res.Time.Hour()
	if hour < openHour || hour >= closeHour {
		http.Error(w, "Reservation time outside operating hours", http.StatusBadRequest)
		return
	}

	dateKey := res.Time.Format("2006-01-02 15:00")
	db.Lock()
	defer db.Unlock()

	if len(db.reservations[dateKey]) >= db.reservationCapacity {
		http.Error(w, "Time slot is full", http.StatusConflict)
		return
	}

	db.reservations[dateKey] = append(db.reservations[dateKey], res)
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Reservation confirmed for %s at %s", res.Name, res.Time.Format("3:04 PM"))
}

// viewReservations - View existing reservation ?date=2024-08-21
func viewReservations(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting viewReservations()")

	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}
	log.Printf("Fetch reservations for %s", dateStr)

	dateKeyPrefix := date.Format("2006-01-02")
	db.Lock()
	defer db.Unlock()

	var dayReservations []Reservation
	for timeSlot, resList := range db.reservations {
		if len(timeSlot) >= len(dateKeyPrefix) && timeSlot[:len(dateKeyPrefix)] == dateKeyPrefix {
			dayReservations = append(dayReservations, resList...)
		}
	}

	if len(dayReservations) == 0 {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, "No reservations found for the given date")
		return
	}

	_ = json.NewEncoder(w).Encode(dayReservations)
}

// cancelReservation - delete an existing reservation
func cancelReservation(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting cancelReservation()")

	var res Reservation
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dateKey := res.Time.Format("2006-01-02 15:00")
	db.Lock()
	defer db.Unlock()

	if resList, ok := db.reservations[dateKey]; ok {
		for i, existingRes := range resList {
			if existingRes.Name == res.Name {
				db.reservations[dateKey] = append(resList[:i], resList[i+1:]...)
				w.WriteHeader(http.StatusOK)
				_, _ = fmt.Fprintf(w, "Reservation canceled for %s at %s",
					res.Name, res.Time.Format("3:04 PM"))
				return
			}
		}
	}

	http.Error(w, "Reservation not found", http.StatusNotFound)
}

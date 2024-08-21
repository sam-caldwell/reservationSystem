package data

import (
	"fmt"
	"log"
)

// AddReservation - add a reservation to our data store
func (db *Database) AddReservation(reservation *Reservation) (err error, response *ReservationResponse) {
	log.Println("AddReservation")
	hour := reservation.Time.Hour()
	dateKey := reservation.Time.Format("2006-01-02 15")

	response = &ReservationResponse{
		Time: dateKey,
		Name: reservation.Name,
	}

	if hour < db.OpenHour || hour >= db.CloseHour {
		response.Status = "Reservation time outside of normal operating hours"
		return nil, response
	}

	db.Lock()
	defer db.Unlock()

	if len(db.reservations[dateKey]) >= db.reservationCapacity {
		response.Status = "Time slot is full (we are adding you to our waitlist"
		db.waitlist=append(db.waitlist,*reservation)
		return nil, response
	}

	db.reservations[dateKey] = append(db.reservations[dateKey], *reservation)
	response.Status = fmt.Sprintf("Reservation confirmed for %s at %s:00", reservation.Name, dateKey)
	return nil, response
}

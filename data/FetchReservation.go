package data

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// FetchReservations - fetch some set of reservations from our data store
func (db *Database) FetchReservations(date time.Time) (err error, resultSet *ReservationSet) {
	log.Println("FetchReservations()")
	db.Lock()
	defer db.Unlock()

	//Day under query
	reservationDay := date.Format("2006-01-02")
	resultSet = &ReservationSet{}

	for timeSlot, reservationList := range db.reservations {
		log.Printf("  dump - %v: %v", timeSlot, reservationList)

		//Gather records for the current day
		if thisDay := strings.Split(timeSlot, " ")[0]; thisDay == reservationDay {
			for _, reservation := range reservationList {
				resultSet.List = append(resultSet.List,
					fmt.Sprintf("%s : %s", reservation.Time, reservation.Name))
			}
		}
	}
	return err, resultSet
}

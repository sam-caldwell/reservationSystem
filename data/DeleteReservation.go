package data

import (
	"fmt"
)

// DeleteReservation - delete a reservation from our data store
func (db *Database) DeleteReservation(reservation *Reservation) (err error, msg string) {
	db.Lock()
	defer db.Unlock()

	dateKey := reservation.Time.Format("2006-01-02 15")

	if resList, ok := Db.reservations[dateKey]; ok {
		for i, existingRes := range resList {
			if existingRes.Name == reservation.Name {
				db.reservations[dateKey] = append(resList[:i], resList[i+1:]...)
				return nil, fmt.Sprintf("Reservation canceled for %s at %s",
					reservation.Name, dateKey)
			}
		}
	}

	return nil, fmt.Sprintf("Reservation not found for %s at %s", reservation.Name, dateKey)

}

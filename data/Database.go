package data

import "sync"

// Database - Database Schema
type Database struct {
	sync.Mutex

	OpenHour int

	CloseHour int

	reservationCapacity int

	reservations map[string][]Reservation

	waitlist []Reservation
}

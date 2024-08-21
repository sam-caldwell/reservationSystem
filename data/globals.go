package data

const (
	defaultCapacity  = 4
	defaultOpenHour  = 10
	defaultCloseHour = 22
)

// Db - a fake in-memory database for our reservations
var Db = Database{
	reservationCapacity: defaultCapacity,
	OpenHour:            defaultOpenHour,
	CloseHour:           defaultCloseHour,
	reservations:        make(map[string][]Reservation),
	waitlist:            []Reservation),
}

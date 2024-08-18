package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Reservation - a single reservation entity
type Reservation struct {
	Name string    `json:"name"`
	Time time.Time `json:"time"` //Note: this is going to expect  RFC 3339 time (in zulu)
}

// Unmarshall - convert http bytes request to Reservation struct
func (req *Reservation) Unmarshall(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return fmt.Errorf("error decoding json: %v", err)
	}
	//Make sure all our reservation times are in even hours
	req.Time = time.Date(
		req.Time.Year(),
		req.Time.Month(),
		req.Time.Day(),
		req.Time.Hour(),
		0, 0, 0, time.Local)
	return nil
}

package data

import (
	"encoding/json"
	"fmt"
)

// ReservationResponse - a response to a reservation operation (add/delete)
type ReservationResponse struct {
	Name   string `json:"name"`
	Time   string `json:"time"`
	Status string `json:"status"`
}

// Marshal - Marshall the ReservationResponse into a json byte string
func (resultSet *ReservationResponse) Marshal() (error, []byte) {
	jsonData, err := json.Marshal(resultSet)
	if err != nil {
		return fmt.Errorf("error marshalling reservation: %v", err), nil
	}
	return nil, jsonData
}

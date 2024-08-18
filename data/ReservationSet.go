package data

import (
	"encoding/json"
	"fmt"
)

// ReservationSet - a set of reservations in response to a query
type ReservationSet struct {
	Count int      `json:"count"`
	List  []string `json:"list"`
}

// Marshal - Marshall the ReservationSet into a json byte string
func (resultSet *ReservationSet) Marshal() (error, []byte) {
	jsonData, err := json.Marshal(resultSet)
	if err != nil {
		return fmt.Errorf("error marshalling reservation: %v", err), nil
	}
	return nil, jsonData
}

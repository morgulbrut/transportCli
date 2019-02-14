package parseJSON

import (
	"bytes"
	"encoding/json"
)

type Coordinate struct {
	Type string
	X    float64
	Y    float64
}

type Station struct {
	ID         string
	Name       string
	Score      string
	Distance   float64
	Coordinate Coordinate
}

type Prognosis struct {
	Platform    string
	Arrival     string
	Departure   string
	Capacity1st string
	Capacity2nd string
}

type Stop struct {
	Station            Station
	Arrival            string
	ArrivalTimestamp   int
	Departure          string
	DepartureTimestamp int
	Delay              string
	Platform           string
	Prognosis          Prognosis
}

func ParseJSON(data []byte) {
	json.MarshalIndent(data, "", "    ")
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

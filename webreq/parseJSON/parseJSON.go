package parseJSON

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	Station              Station
	Arrival              string
	ArrivalTimestamp     int64
	Departure            string
	DepartureTimestamp   int64
	Delay                string
	Platform             string
	Prognosis            Prognosis
	RealtimeAvailability string
	Location             Station
	Distance             float64
}

type Stationboard struct {
	Stop         Stop
	Name         string
	Category     string
	Subcategory  string
	CategoryCode string
	Number       string
	Operator     string
	To           string
	Capacity1st  string
	Capacity2nd  string
	Passlist     []Stop
}

type ResponseStation struct {
	Station      Station
	Stationboard []Stationboard
}

func ParseStation(data []byte) ResponseStation {
	var resp ResponseStation

	err := json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Printf("%+v", resp)
	return resp
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

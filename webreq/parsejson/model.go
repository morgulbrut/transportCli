/*
Copyright © 2019 morgulbrut
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file or
 http://www.wtfpl.net/ for more details.
*/

package parsejson

type RespConnection struct {
	Connections []Connections `json:"connections"`
	From        Stop          `json:"from"`
	To          Stop          `json:"to"`
	Stations    Stations      `json:"stations"`
}

type Journey struct {
	Name         string      `json:"name"`
	Category     string      `json:"category"`
	Subcategory  interface{} `json:"subcategory"`
	CategoryCode interface{} `json:"categoryCode"`
	Number       string      `json:"number"`
	Operator     string      `json:"operator"`
	To           string      `json:"to"`
	PassList     []PassList  `json:"passList"`
	Capacity1St  interface{} `json:"capacity1st"`
	Capacity2Nd  interface{} `json:"capacity2nd"`
}
type Departure struct {
	Station              Station     `json:"station"`
	Arrival              interface{} `json:"arrival"`
	ArrivalTimestamp     interface{} `json:"arrivalTimestamp"`
	Departure            string      `json:"departure"`
	DepartureTimestamp   int         `json:"departureTimestamp"`
	Delay                int         `json:"delay"`
	Platform             string      `json:"platform"`
	Prognosis            Prognosis   `json:"prognosis"`
	RealtimeAvailability interface{} `json:"realtimeAvailability"`
	Location             Location    `json:"location"`
}
type Arrival struct {
	Station              Station     `json:"station"`
	Arrival              string      `json:"arrival"`
	ArrivalTimestamp     int         `json:"arrivalTimestamp"`
	Departure            interface{} `json:"departure"`
	DepartureTimestamp   interface{} `json:"departureTimestamp"`
	Delay                int         `json:"delay"`
	Platform             string      `json:"platform"`
	Prognosis            Prognosis   `json:"prognosis"`
	RealtimeAvailability interface{} `json:"realtimeAvailability"`
	Location             Location    `json:"location"`
}
type Sections struct {
	Journey   Journey     `json:"journey"`
	Walk      interface{} `json:"walk"`
	Departure Departure   `json:"departure"`
	Arrival   Arrival     `json:"arrival"`
}
type Connections struct {
	From        Stop        `json:"from"`
	To          Stop        `json:"to"`
	Duration    string      `json:"duration"`
	Transfers   int         `json:"transfers"`
	Service     interface{} `json:"service"`
	Products    []string    `json:"products"`
	Capacity1St interface{} `json:"capacity1st"`
	Capacity2Nd interface{} `json:"capacity2nd"`
	Sections    []Sections  `json:"sections"`
}
type From struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Score      interface{} `json:"score"`
	Coordinate Coordinate  `json:"coordinate"`
	Distance   interface{} `json:"distance"`
}
type To struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Score      interface{} `json:"score"`
	Coordinate Coordinate  `json:"coordinate"`
	Distance   interface{} `json:"distance"`
}
type Stations struct {
	From []From `json:"from"`
	To   []To   `json:"to"`
}

type RespStation struct {
	Station      Station        `json:"station"`
	Stationboard []Stationboard `json:"stationboard"`
}
type Coordinate struct {
	Type string  `json:"type"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
}
type Station struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Score      interface{} `json:"score"`
	Coordinate Coordinate  `json:"coordinate"`
	Distance   interface{} `json:"distance"`
}
type Prognosis struct {
	Platform    interface{} `json:"platform"`
	Arrival     interface{} `json:"arrival"`
	Departure   interface{} `json:"departure"`
	Capacity1St interface{} `json:"capacity1st"`
	Capacity2Nd interface{} `json:"capacity2nd"`
}
type Location struct {
	ID         string      `json:"id"`
	Name       interface{} `json:"name"`
	Score      interface{} `json:"score"`
	Coordinate Coordinate  `json:"coordinate"`
	Distance   interface{} `json:"distance"`
}
type Stop struct {
	Station              Station     `json:"station"`
	Arrival              interface{} `json:"arrival"`
	ArrivalTimestamp     interface{} `json:"arrivalTimestamp"`
	Departure            string      `json:"departure"`
	DepartureTimestamp   int         `json:"departureTimestamp"`
	Delay                interface{} `json:"delay"`
	Platform             string      `json:"platform"`
	Prognosis            Prognosis   `json:"prognosis"`
	RealtimeAvailability interface{} `json:"realtimeAvailability"`
	Location             Location    `json:"location"`
}
type PassList struct {
	Station              Station     `json:"station"`
	Arrival              interface{} `json:"arrival"`
	ArrivalTimestamp     interface{} `json:"arrivalTimestamp"`
	Departure            string      `json:"departure"`
	DepartureTimestamp   int         `json:"departureTimestamp"`
	Delay                interface{} `json:"delay"`
	Platform             string      `json:"platform"`
	Prognosis            Prognosis   `json:"prognosis"`
	RealtimeAvailability interface{} `json:"realtimeAvailability"`
	Location             Location    `json:"location"`
}
type Stationboard struct {
	Stop         Stop        `json:"stop"`
	Name         string      `json:"name"`
	Category     string      `json:"category"`
	Subcategory  interface{} `json:"subcategory"`
	CategoryCode interface{} `json:"categoryCode"`
	Number       string      `json:"number"`
	Operator     string      `json:"operator"`
	To           string      `json:"to"`
	PassList     []PassList  `json:"passList"`
	Capacity1St  interface{} `json:"capacity1st"`
	Capacity2Nd  interface{} `json:"capacity2nd"`
}

type RespLocation struct {
	Stations []Stations `json:"stations"`
}

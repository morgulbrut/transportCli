/*
Copyright © 2019 morgulbrut
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file or
 http://www.wtfpl.net/ for more details.
*/

package parsejson

// Location represents a API location object
type Location struct {
	ID          string
	Type        string
	Name        string
	Score       string
	Coordinates Coordinate
	Distance    int
}

// Coordinate represents a API coordinate object
type Coordinate struct {
	Type string
	X    float64
	Y    float64
}

// Connection represents a API connection object
type Connection struct {
	From        Stop
	To          Stop
	Duration    string
	Service     Service
	Products    []string
	Capacity1st string
	Capacity2nd string
}

// Service represents a API service object
type Service struct {
	Regular   string
	Irregular string
}

// Prognosis represents a API prognosis object
type Prognosis struct {
	Platform    string
	Arrival     string
	Departure   string
	Capacity1st string
	Capacity2nd string
}

// Stop represents a API stop object
type Stop struct {
	Station   Location
	Arrival   string
	Departure string
	Delay     string
	Platform  string
	Prognosis Prognosis
}

// Section represents a API section object
type Section struct {
	Journey   Journey
	Walk      string
	Departure Stop
	Arrival   Stop
}

// Journey represents a API journey object
type Journey struct {
	Name         string
	Category     string
	CategoryCode string
	Number       string
	Operator     string
	To           string
	PassList     []Stop
	Capacity1st  string
	Capacity2nd  string
}

// RespStation represents a response of a station API call
type RespStation struct {
	Station      Location
	Stationboard []Journey
}

// RespLocation represents a response of a location API call
type RespLocation struct {
	Locations []Location
}

// RespConnection represents a response of a connection API call
type RespConnection struct {
	Connections []Connection
}

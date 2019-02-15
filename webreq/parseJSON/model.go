/*
Copyright © 2019 morgulbrut
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file or
 http://www.wtfpl.net/ for more details.
*/

package parseJSON

type Location struct {
	ID          string
	Type        string
	Name        string
	Score       string
	Coordinates Coordinate
	Distance    int
}

type Coordinate struct {
	Type string
	X    float64
	Y    float64
}

type Connection struct {
	From        Stop
	To          Stop
	Duration    string
	Service     Service
	Products    []string
	Capacity1st string
	Capacity2nd string
}

type Service struct {
	Regular   string
	Irregular string
}

type Prognosis struct {
	Platform    string
	Arrival     string
	Departure   string
	Capacity1st string
	Capacity2nd string
}

type Stop struct {
	Station   Location
	Arrival   string
	Departure string
	Delay     string
	Platform  string
	Prognosis Prognosis
}

type Section struct {
	Journey   Journey
	Walk      string
	Departure Stop
	Arrival   Stop
}

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

type RespStation struct {
	Station      Location
	Stationboard []Journey
}

type RespLocation struct {
	Locations []Location
}

type RespConnection struct {
	Connections []Connection
}

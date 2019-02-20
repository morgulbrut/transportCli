/*
Copyright © 2019 morgulbrut
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file or
 http://www.wtfpl.net/ for more details.
*/

package webreq

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/morgulbrut/transportCli/webreq/parsejson"
)

const baseURL string = "http://transport.opendata.ch"
const stationURL string = "/v1/stationboard"
const locationURL string = "/v1/locations"

// Station does the API call for a station and returns a RespStation object
func Connection(args string) parsejson.RespConnection {
	body := webreq(stationURL, args)
	return parsejson.ParseConnection(body)
}

// Station does the API call for a station and returns a RespStation object
func Station(args string) parsejson.RespStation {
	body := webreq(stationURL, args)
	return parsejson.ParseStation(body)
}

// Location does the API call for a location and returns a RespLocation object
func Location(args string) parsejson.RespLocation {
	body := webreq(locationURL, args)
	return parsejson.ParseLocation(body)
}

func webreq(resourceURL string, args string) []byte {
	a := strings.Replace(args, " ", "%20", -1)
	wr := baseURL + resourceURL + a
	fmt.Printf("GET %s\n", wr)
	resp, err := http.Get(wr)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

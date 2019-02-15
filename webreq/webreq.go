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

	"github.com/morgulbrut/transportCli/webreq/parseJSON"
)

const baseURL string = "http://transport.opendata.ch"
const stationURL string = "/v1/stationboard"
const locationURL string = "/v1/locations"

func Station(args string) parseJSON.RespStation {
	body := webreq(stationURL, args)
	return parseJSON.ParseStation(body)
}

func Location(args string) parseJSON.RespLocation {
	body := webreq(locationURL, args)
	return parseJSON.ParseLocation(body)
}

func webreq(resourceURL string, args string) []byte {
	wr := baseURL + resourceURL + args
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

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

const BaseURL string = "http://transport.opendata.ch"
const stationURL string = "/v1/stationboard"

func WebreqStation(args string) parseJSON.ResponseStation {
	body := webreq(stationURL, args)
	return parseJSON.ParseStation(body)
}

func webreq(resourceURL string, args string) []byte {
	resp, err := http.Get(BaseURL + resourceURL + args)
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

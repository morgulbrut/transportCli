package webreq

import (
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

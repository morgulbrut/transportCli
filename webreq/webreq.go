package webreq

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const BaseURL string = "http://transport.opendata.ch"

func Webreq(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

}

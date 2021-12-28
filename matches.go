package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Matches fetches matches by id.
func Matches(id int) ([]byte, error) {

	url := fmt.Sprintf("https://api.opendota.com/api/matches/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Makes a GET call to the API.
func Get(path string, params *url.Values) ([]byte, error) {

	base := fmt.Sprintf("%s%s", URL, path)

	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	if params != nil && len(*params) > 0 {
		u.RawQuery = params.Encode()
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

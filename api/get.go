package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Makes a GET call to the API.
func Get(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", URL, path)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

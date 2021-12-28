package main

import (
	"fmt"
	"net/url"
)

const (
	SearchPath = "search"
)

// Search fetches matches by id.
func Search(name string) ([]byte, error) {
	path := fmt.Sprintf("/%s", SearchPath)

	params := url.Values{}
	params.Add("q", name)

	return Get(path, params)
}

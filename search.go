package main

import "fmt"

// Search fetches matches by id.
func Search(name string) ([]byte, error) {
	url := fmt.Sprintf("/search/%d", name)

	return Get(url)
}

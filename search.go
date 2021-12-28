package main

import "fmt"

// Search fetches matches by id.
func Search(name string) ([]byte, error) {
	path := fmt.Sprintf("/search/%d", name)

	return Get(path, nil)
}

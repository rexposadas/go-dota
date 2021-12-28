package main

import "fmt"

// Player fetches matches by id.
func Player(id int) ([]byte, error) {
	path := fmt.Sprintf("/players/%d", id)

	return Get(path, nil)
}

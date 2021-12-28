package main

import (
	"fmt"
)

// Match fetches matches by id.
func Match(id int) ([]byte, error) {
	path := fmt.Sprintf("/matches/%d", id)

	return Get(path, nil)
}

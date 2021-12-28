package main

import (
	"api"
	"fmt"
)

// Match fetches matches by id.
func Match(id int) ([]byte, error) {
	url := fmt.Sprintf("/matches/%d", id)

	return api.Get(url)
}

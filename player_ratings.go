package main

import "fmt"

// PlayerRating fetches matches by id.
func PlayerRating(id int) ([]byte, error) {
	url := fmt.Sprintf("/players/%d/ratings", id)

	return Get(url)
}

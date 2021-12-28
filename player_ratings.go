package main

import "fmt"

// PlayerRating fetches matches by id.
func PlayerRating(id int) ([]byte, error) {
	path := fmt.Sprintf("/players/%d/ratings", id)

	return Get(path, nil )
}

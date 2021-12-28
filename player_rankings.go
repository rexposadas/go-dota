package main

import "fmt"

// PlayerRating fetches matches by id.
func PlayerRankings(id int) ([]byte, error) {
	path := fmt.Sprintf("/players/%d/rankings", id)

	return Get(path, nil)
}

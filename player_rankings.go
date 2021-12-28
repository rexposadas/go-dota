package main

import "fmt"

// PlayerRating fetches matches by id.
func PlayerRankings(id int) ([]byte, error) {
	url := fmt.Sprintf("/players/%d/rankings", id)

	return Get(url)
}

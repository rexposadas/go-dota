package main

import "fmt"

// ItemPopularity fetches matches by id.
func ItemPopularity(id int) ([]byte, error) {
	url := fmt.Sprintf("/heroes/%d/itemPopularity", id)
	return Get(url)
}

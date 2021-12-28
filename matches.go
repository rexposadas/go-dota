package main

import "fmt"

func Match(id int) ([]byte, error) {
	url := fmt.Sprintf("/matches/%d", id)

	return Get(url)
}

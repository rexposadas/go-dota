package main

import (
	"log"
)

func main() {

	// b, err := ProPlayers()
	// params := url.Values{}
	// params.Add("limit", "10")

	b, err := Search("Gabbi")
	if err != nil {
		log.Printf("%s", err)
	}

	log.Printf("results: %s", string(b))
}

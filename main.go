package main

import "log"

func main() {

	b, err := PlayerRating(226274743)
	if err != nil {
		log.Printf("%s", err)
	}

	log.Printf("matches %s", string(b))
}

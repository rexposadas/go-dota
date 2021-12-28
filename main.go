package main

import "log"

func main() {

	// b, err := ProPlayers(226274743)
	b, err := PublicMatches()
	if err != nil {
		log.Printf("%s", err)
	}

	log.Printf("matches %s", string(b))
}

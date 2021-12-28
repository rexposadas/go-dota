package main

import "log"

func main() {

	b, err := Player(226274743)
	if err != nil {
		log.Printf("%s", err)
	}

	log.Printf("matches %s", string(b))
}

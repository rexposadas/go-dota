package main

import "log"

func main() {

	b, err := Matches(271145478)
	if err != nil {
		log.Printf("%s", err)
	}

	log.Printf("matches %s", string(b))
}

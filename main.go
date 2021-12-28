package main

import "log"


func main() {

	b, err := PlayersByRank()
	if err != nil {
		log.Printf("%s", err)
	}

	log.Printf("matches %s", string(b))
}

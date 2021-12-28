package main

import "fmt"

// Heroes fetches matches by id.
func Heroes() ([]byte, error) {
	return Get("/heroes")
}

// See:https://docs.opendota.com/#tag/heroes%2Fpaths%2F~1heroes~1%7Bhero_id%7D~1matches%2Fget
func MatchesByHeroID(id int) ([]byte, error) {
	path := fmt.Sprintf("/heroes/%d/matches", id)
	return Get(path)
}

func HeroMatchups(id int) ([]byte, error) {
	path := fmt.Sprintf("/heroes/%d/matchups", id)
	return Get(path)
}

func HeroDuration(id int) ([]byte, error) {
	path := fmt.Sprintf("/heroes/%d/durations", id)
	return Get(path)
}

// Get players who have played this hero
// See: https://docs.opendota.com/#tag/heroes%2Fpaths%2F~1heroes~1%7Bhero_id%7D~1players%2Fget
func PlayersWhoPlayedHero(id int) ([]byte, error) {
	path := fmt.Sprintf("/heroes/%d/players", id)
	return Get(path)
}

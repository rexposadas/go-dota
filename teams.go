package main

import "fmt"

const (
	TEAMS = "teams"
)

// Teams fetches matches by id.
func Teams() ([]byte, error) {
	return Get("/" + TEAMS)
}

// Get data for a Team
func TeamData(id int) ([]byte, error) {
	path := fmt.Sprintf("/%s/%d", TEAMS, id)
	return Get(path)
}

// Get matches for a team.
func TeamMatches(id int) ([]byte, error) {
	path := fmt.Sprintf("/%s/%d/matches", TEAMS, id)
	return Get(path)
}

func TeamPlayers(id int) ([]byte, error) {
	path := fmt.Sprintf("%s/%d/players", TEAMS, id)
	return Get(path)
}

func TeamHeroes(id int) ([]byte, error) {
	path := fmt.Sprintf("%s/%d/heroes", TEAMS, id)
	return Get(path)
}

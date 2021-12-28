package main

import "fmt"

const (
	LEAGUES = "leagues"
)

// Leagues fetches matches by id.
func Leagues() ([]byte, error) {
	return Get("/"+LEAGUES, nil)
}

// Get data for a league
func LeagueData(id int) ([]byte, error) {
	path := fmt.Sprintf("/%s/%d", LEAGUES, id)
	return Get(path, nil)
}

// Get matches for a team.
func LeagueMatches(id int) ([]byte, error) {
	path := fmt.Sprintf("/%s/%d/matches", LEAGUES, id)
	return Get(path, nil)
}

func LeagueTeams(id int) ([]byte, error) {
	path := fmt.Sprintf("%s/%d/teams", LEAGUES, id)
	return Get(path, nil)
}

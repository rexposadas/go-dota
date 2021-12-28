package main

import "fmt"

const (
	LEAGUES = "leagues"
)

// Leagues fetches matches by id.
func Leagues() ([]byte, error) {
	return Get("/" + LEAGUES)
}

// Get data for a league
func LeagueData(id int) ([]byte, error) {
	path := fmt.Sprintf("/%s/%d", LEAGUES, id)
	return Get(path)
}

// Get matches for a team.
func LeagueMatches(id int) ([]byte, error) {
	path := fmt.Sprintf("/%s/%d/matches", LEAGUES, id)
	return Get(path)
}

func LeagueTeams(id int) ([]byte, error) {
	path := fmt.Sprintf("%s/%d/teams", LEAGUES, id)
	return Get(path)
}

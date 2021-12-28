package main

import (
	"fmt"
	"net/url"
)

const (
	PlayersBasePath = "players"
)

// Player fetches matches by id.
func Player(id int) ([]byte, error) {
	path := fmt.Sprintf("/%s/%d", PlayersBasePath, id)

	return Get(path, nil)
}

func PlayerRecentMatches(id int) ([]byte, error) {
	path := fmt.Sprintf("/%s/%d/recentMatches", PlayersBasePath, id)

	return Get(path, nil)
}

func PlayerWL(id int, params url.Values) ([]byte, error) {

	path := fmt.Sprintf("/%s/%d/wl", PlayersBasePath, id)

	return Get(path, params)
}

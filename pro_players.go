package main

// ProPlayers fetches matches by id.
func ProPlayers() ([]byte, error) {
	return Get("/proPlayers", nil )
}

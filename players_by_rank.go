package main

// PlayersByRank fetches matches by id.
func PlayersByRank() ([]byte, error) {
	return Get("/playersByRank")
}

package main

// ProMatches fetches matches by id.
func ProMatches() ([]byte, error) {
	return Get("/proMatches", nil)
}

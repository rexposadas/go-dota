package main

// PublicMatches fetches matches by id.
func PublicMatches() ([]byte, error) {
	return Get("/proMatches", nil)
}

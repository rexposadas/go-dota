package main

// ParsedMatches fetches matches by id.
func ParsedMatches() ([]byte, error) {
	return Get("/parsedMatches", nil )
}

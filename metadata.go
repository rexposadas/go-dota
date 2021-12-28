package main

// Metadata fetches matches by id.
func Metadata() ([]byte, error) {
	return Get("/metadata")
}

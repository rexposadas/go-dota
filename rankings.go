package main

// Rankings fetches matches by id.
func Rankings() ([]byte, error) {
	return Get("/rankings", nil)
}

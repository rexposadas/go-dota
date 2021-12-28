package main

// Health fetches matches by id.
func Health() ([]byte, error) {
	return Get("/health")
}

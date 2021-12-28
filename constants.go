package main

// Constants fetches matches by id.
func Constants() ([]byte, error) {
	return Get("/constants", nil )
}

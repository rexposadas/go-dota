package main

// Distributions fetches matches by id.
func Distributions() ([]byte, error) {
	return Get("/distributions")
}

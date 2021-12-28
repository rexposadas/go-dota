package main

// Status fetches matches by id.
func Status() ([]byte, error) {
	return Get("/status")
}

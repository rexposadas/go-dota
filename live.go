package main

// Live fetches matches by id.
func Live() ([]byte, error) {
	return Get("/live")
}

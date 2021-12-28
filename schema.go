package main

// Schema fetches matches by id.
func Schema() ([]byte, error) {
	return Get("/schema", nil )
}

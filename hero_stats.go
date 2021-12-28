package main

// HeroStats fetches matches by id.
func HeroStats() ([]byte, error) {
	return Get("/heroStats")
}

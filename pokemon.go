package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Pokemon struct maps only the fields we need
type Pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}

// FetchPokemon retrieves data from the PokeAPI and returns a Pokemon struct
func FetchPokemon(name string) (Pokemon, error) {
	var p Pokemon

	// Build URL
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

	// Use a custom HTTP client with timeout
	client := &http.Client{Timeout: 5 * time.Second}

	// Send GET request
	resp, err := client.Get(url)
	if err != nil {
		return p, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	// Check HTTP status
	if resp.StatusCode != http.StatusOK {
		return p, fmt.Errorf("Pokémon '%s' not found (status %d)", name, resp.StatusCode)
	}

	// Decode JSON into struct
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return p, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return p, nil
}

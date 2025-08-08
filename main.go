package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if argument containing pokemon name exists
	if len(os.Args) < 2 {
		fmt.Println("Usage: pokedex-cli <pokemon-name>")
		return
	}

	pokemonName := strings.ToLower(os.Args[1])

	p, err := FetchPokemon(pokemonName)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	PrintPokemon(p)

	//fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", p.Name, p.Height, p.Weight)

}

// Things to take note of:

func (p Pokemon) Summary() string { // value receiver - works on a copy
	return fmt.Sprintf("%s: h=%d w=%d", p.Name, p.Height, p.Weight)
}

func (p *Pokemon) SetName(n string) { // pointer receiver - modifies original
	p.Name = n
}

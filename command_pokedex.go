package main

import "fmt"

func commandPokedex(c *config, args ...string) error {
	println("Your Pokedex:")
	for _, pok := range c.caughtPokemon {
		fmt.Printf(" - %s\n", pok.Name)
	}
	return nil
}

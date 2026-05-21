package main

import "fmt"

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println()
	return nil
}

package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a captured pokemon name")
	}
	name := args[0]
	caughtpokemon, ok := c.caughtPokemon[name]
	if !ok {
		return errors.New("pokemon hasn't been caught")
	}
	fmt.Println("Name:", caughtpokemon.Name)
	fmt.Println("Height:", caughtpokemon.Height)
	fmt.Println("Weight:", caughtpokemon.Weight)
	fmt.Println("Stats:")
	for _, pok := range caughtpokemon.Stats {
		fmt.Printf(" -%s: %v\n", pok.Stat.Name, pok.BaseStat)
	}
	fmt.Println("Types:")
	for _, pok := range caughtpokemon.Types {
		fmt.Printf(" -%s\n", pok.Type.Name)
	}

	return nil
}

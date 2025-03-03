package main

import (
	"fmt"
)



func commandInspect(con *config, pokeName string) error {

	pokedexEntry, ok := con.Pokedex.Get(pokeName)
	if !ok {
		return fmt.Errorf("%s pokemon is not in your pokedex", pokeName)
	}
	fmt.Printf("Name: %s\n",pokedexEntry.Name)
	fmt.Printf("Height: %d\n",pokedexEntry.Height)
	fmt.Printf("Weight: %d\n",pokedexEntry.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokedexEntry.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pokeType := range pokedexEntry.Types {
		fmt.Printf(" -%s\n", pokeType.Type.Name)
	}
	return nil
}
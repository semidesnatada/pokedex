package main

import "fmt"

func commandPokedex(con *config, parameter string) error {

	fmt.Println("Your Pokedex:")
	for pokemon, _ := range con.Pokedex.Content {
		fmt.Println(" - ", pokemon)
	}
	return nil

}
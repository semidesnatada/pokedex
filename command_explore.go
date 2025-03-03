package main

import "fmt"


func commandExplore(con *config, specificLocation string) error {

	// fmt.Println("used explore command in area: ", specificLocation)

	data, err := con.PokeClient.GetLocationSpecificData(specificLocation)
	if err != nil {
		return err
	}

	fmt.Println("Exploring", specificLocation, "...")
	fmt.Println("Found Pokemon:")
	for _, item := range data.PokemonEncounters {
		fmt.Println(" - ",item.Pokemon.Name)
	}

	return nil

}
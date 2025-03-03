package main

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/semidesnatada/pokedex/internal/pokecache"
)


func commandCatch(con *config, pokeName string) error {

	data, err := con.PokeClient.GetPokemonResponseData(pokeName)
	if err != nil {
		return err
	}

	//use baseExperience as a parameter to be used to set catch probabiliy
	baseExp := data.BaseExperience
	var catchChance float64
	if baseExp < 300 {
		catchChance = float64(baseExp) / 750.0 
	} else {
		catchChance = math.Pow(float64(1) - float64(baseExp)/1000.0, 5)
	}
	catchChance += 0.2
	fmt.Printf("Chance of catching %s is %.2f\n", pokeName, catchChance)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)

	threshold := rand.Float64()
	if catchChance > threshold {
		fmt.Println(pokeName, " was caught!")
		con.Pokedex.Add(pokecache.PokedexEntry{
			Name: pokeName,
			Height: data.Height,
			Weight: data.Weight,
			Stats: data.Stats,
			Types: data.Types,
		})
	} else {
		fmt.Println(pokeName, " escaped!")
	}

	return nil


}
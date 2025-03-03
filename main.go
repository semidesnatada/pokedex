package main

import (
	"time"

	"github.com/semidesnatada/pokedex/internal/pokeapi"
	"github.com/semidesnatada/pokedex/internal/pokecache"
)


func main() {

	client := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(7 * time.Second)
	pokedex := pokecache.NewPokedex()
	con := config{
		PokeClient: client,
		Pokedex: pokedex,
		Cache: cache,
	}

	startRepl(&con)

}

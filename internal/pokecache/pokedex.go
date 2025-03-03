package pokecache

import (
	"github.com/semidesnatada/pokedex/internal/pokeapi"
)

type Pokedex struct {
	Content map[string]PokedexEntry
}

type PokedexEntry struct {
	Name string
	Height int
	Weight int
	Stats pokeapi.StatsArray
	Types pokeapi.TypesArray
}

func NewPokedex() Pokedex {
	return Pokedex{
		Content: map[string]PokedexEntry{},
	}
}

func (p *Pokedex) Add(pokemon PokedexEntry) {
	p.Content[pokemon.Name] = pokemon
}

func (p *Pokedex) Get(pokemonName string) (PokedexEntry, bool) {

	item, ok := p.Content[pokemonName]
	if !ok {
		return PokedexEntry{}, false
	}
	return item, true
}
package pokecache

import (
	"fmt"
	"testing"
	"time"

	"github.com/semidesnatada/pokedex/internal/pokeapi"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val pokeapi.LocationListResponseFormat
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area/",
			val: pokeapi.LocationListResponseFormat{
				Count: 1089,
				Next: "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
				Previous: "",
				Results: []pokeapi.NameUrlPair{
					{
					Name: "canalave-city-area",
					Url: "https://pokeapi.co/api/v2/location-area/1/",
					},
					{
					Name: "eterna-city-area",
					Url: "https://pokeapi.co/api/v2/location-area/2/",
					},
					{
						Name: "pastoria-city-area",
						Url: "https://pokeapi.co/api/v2/location-area/3/",
					},
					{
						Name:  "sunyshore-city-area",
						Url: "https://pokeapi.co/api/v2/location-area/4/",
					},	
					{
						Name: "sinnoh-pokemon-league-area",
						Url: "https://pokeapi.co/api/v2/location-area/5/",
					},
					{
						Name: "oreburgh-mine-1f",
						Url: "https://pokeapi.co/api/v2/location-area/6/",
					},
					{
						Name: "oreburgh-mine-b1f",
						Url: "https://pokeapi.co/api/v2/location-area/7/",
					},
					{
						Name:"valley-windworks-area",
						Url: "https://pokeapi.co/api/v2/location-area/8/",
					},
					{
						Name: "eterna-forest-area",
						Url: "https://pokeapi.co/api/v2/location-area/9/",
					},
					{
						Name:"fuego-ironworks-area",
						Url: "https://pokeapi.co/api/v2/location-area/10/",
					},
					{
						Name:"mt-coronet-1f-route-207",
						Url: "https://pokeapi.co/api/v2/location-area/11/",
					},
					{
						Name:"mt-coronet-2f",
						Url: "https://pokeapi.co/api/v2/location-area/12/",
					},
					{
						Name: "mt-coronet-3f",
						Url: "https://pokeapi.co/api/v2/location-area/13/",
					},
					{
						Name:"mt-coronet-exterior-snowfall",
						Url: "https://pokeapi.co/api/v2/location-area/14/",
					},
					{
						Name:"mt-coronet-exterior-blizzard",
						Url: "https://pokeapi.co/api/v2/location-area/15/",
					},
					{
						Name:"mt-coronet-4f",
						Url: "https://pokeapi.co/api/v2/location-area16/",
					},
					{
						Name:"mt-coronet-4f-small-room",
						Url: "https://pokeapi.co/api/v2/location-area/17/",
					},
					{
						Name:"mt-coronet-5f",
						Url: "https://pokeapi.co/api/v2/location-area/18/",
					},
					{
						Name: "mt-coronet-6f",
						Url: "https://pokeapi.co/api/v2/location-area/19/",
					},
					{
						Name:"mt-coronet-1f-from-exterior",
						Url: "https://pokeapi.co/api/v2/location-area/20/",
					},

				},
	}}}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v",i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if val.Count != c.val.Count || val.Next != c.val.Next {
				t.Errorf("test failed due to incorrect values")
				return
			}
		})
	}
}
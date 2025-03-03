package main

import (
	"errors"
	"fmt"

	"github.com/semidesnatada/pokedex/internal/pokeapi"
)

func commandMap(con *config, parameter string) error {
	
	targetUrl := con.Next

	err := printLocations(targetUrl, con)	
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(con *config, parameter string) error {
	
	targetUrl := con.Previous

	err := printLocations(targetUrl, con)	
	if err != nil {
		return err
	}
	
	return nil
}

func printLocations(targetUrl string, con *config) error {

	if len(targetUrl) == 0 {
		return errors.New("can't go any further")
	}

	var mapData pokeapi.LocationListResponseFormat

	if cachedVal, ok := con.Cache.Get(targetUrl); !ok {
		requestedVal, err := con.PokeClient.GetLocationListData(targetUrl)
		if err != nil {
			return err
		} else {
			mapData = requestedVal
			fmt.Println()
			fmt.Println("Used data from web")
			fmt.Println()
			con.Cache.Add(targetUrl, requestedVal)
		}
	} else {
		mapData = cachedVal
		fmt.Println()
		fmt.Println("Used cached data")
		fmt.Println()
	}

	for _, item := range mapData.Results {
		fmt.Println(item.Name)
	}

	con.Next = mapData.Next
	con.Previous = mapData.Previous

	return nil

}
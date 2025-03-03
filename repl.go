package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/semidesnatada/pokedex/internal/pokeapi"
	"github.com/semidesnatada/pokedex/internal/pokecache"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(con *config) {

	scanner := bufio.NewScanner(os.Stdin)

	con.Next = "https://pokeapi.co/api/v2/location-area/"
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}
		userCommand := text[0]
		var userParam string
		if len(text) > 1 {
			userParam = text[1]
		}
		// fmt.Println("Your command was:", userCommand)
		
		commandsMap := getCommands()
		command, ok := commandsMap[userCommand]
		if ok {
			err := command.callback(con, userParam)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

type cliCommand struct {
	name string
	description string
	callback func(*config, string) error
}

type config struct {
	PokeClient pokeapi.Client
	Pokedex pokecache.Pokedex
	Cache pokecache.Cache
	Next string
	Previous string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
	"exit": {
       	 name:        "exit",
       	 description: "Exit the Pokedex",
       	 callback:    commandExit,
	},
	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
	"map": {
		name: "map",
		description: "Displays a list of locations",
		callback: commandMap,
	},
	"mapb": {
		name: "mapb",
		description: "Displays a list of locations from the previous page",
		callback: commandMapb,
	},
	"explore": {
		name: "explore",
		description: "Shows the pokemon present in a named area",
		callback: commandExplore,
	},
	"catch": {
		name: "catch",
		description: "Attempts to catch a named pokemon",
		callback: commandCatch,
	},
	"pokedex": {
		name: "pokedex",
		description: "Prints the pokemon in your pokedex",
		callback: commandPokedex,
	},
	"inspect": {
		name: "inspect",
		description: "Prints the stats of a named pokemon in your pokedex",
		callback: commandInspect,
	},
	}
}

package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}
		userCommand := text[0]
		fmt.Println("Your command was:", userCommand)
		
		commandsMap := getCommands()
		command, ok := commandsMap[userCommand]
		if ok {
			var con config
			err := command.callback(&con)
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
	callback func(*config) error
}

type config struct {
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
	}
}

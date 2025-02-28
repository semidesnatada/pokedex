package main

import (
	"fmt"
)

func commandHelp(con *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")	
	fmt.Println("Usage:")
	fmt.Println()
	commandsMap := getCommands()
	for _, comm := range commandsMap {
		fmt.Println(comm.name + ": " + comm.description)
	}
	fmt.Println()
	return nil
}

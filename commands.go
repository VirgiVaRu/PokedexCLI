package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

func registerCommands() map[string]cliCommand {

	supportedCommands := map[string]cliCommand{
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},

		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
	}

	return supportedCommands
}


/// Callbacks:

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	supportedCommands := registerCommands()
	for _, command := range supportedCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	
	return nil
}
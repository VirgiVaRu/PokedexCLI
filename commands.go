package main

import (
	"fmt"
	"os"
	"github.com/VirgiVaRu/pokedexcli/internal/PokeAPI"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(*config) error
}

type config struct {
	Next		string
	Previous 	*string
}

func getCommands() map[string]cliCommand {

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

		"map": {
			name:			"map",
			description:	"Displays the name of 20 location areas in the Pokemon world. Each subsequent call displays the next 20 locations, and so on.",
			callback:		commandMap,
		},

		"mapb": {
			name: 			"mapb",
			description: 	"Desplays the previous 20 maps, if possible",
			callback:		commandMapb,
		},
	}

	return supportedCommands
}


/// Callbacks:

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Println()

	return nil
}

func commandMap(config *config) error {
	locationPage := PokeAPI.GetLocationPage(config.Next)

	config.Next = locationPage.Next
	config.Previous = locationPage.Previous
	

	for _, place := range locationPage.Results {
		fmt.Println(place.Name)
	}

	return nil
}

func commandMapb(config *config) error {
	if config.Previous == nil {
		return fmt.Errorf("you're on the first page")
	} else {
		locationPage := PokeAPI.GetLocationPage(*config.Previous)

		config.Next = locationPage.Next
		config.Previous = locationPage.Previous

		for _, place := range locationPage.Results {
			fmt.Println(place.Name)
		}

		return nil
	}
}
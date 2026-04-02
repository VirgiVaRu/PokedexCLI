package main

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/VirgiVaRu/pokedexcli/internal/PokeAPI"
	"github.com/VirgiVaRu/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(*config, pokecache.Cache) error
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

func commandExit(config *config, cache pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, cache pokecache.Cache) error {
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

func commandMap(config *config, cache pokecache.Cache) error {
	val, found := cache.Get(config.Next)
	var locationPage PokeAPI.LocationPage
	if found {
		err := json.Unmarshal(val, &locationPage)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		locationPage = PokeAPI.GetLocationPage(config.Next)
		val, err := json.Marshal(locationPage)
		if err != nil {
			fmt.Println(err)
		}
		cache.Add(config.Next, val)
	}

	config.Next = locationPage.Next
	config.Previous = locationPage.Previous
		
	locationPage.Print()

	return nil
}

func commandMapb(config *config, cache pokecache.Cache) error {
	if config.Previous == nil {
		return fmt.Errorf("you're on the first page")
	}
	val, found := cache.Get(*config.Previous)
	var locationPage PokeAPI.LocationPage
	if found {
		err := json.Unmarshal(val, &locationPage)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		locationPage = PokeAPI.GetLocationPage(*config.Previous)
		val, err := json.Marshal(locationPage)
		if err != nil {
			fmt.Println(err)
		}
		cache.Add(config.Next, val)
	}
	config.Next = locationPage.Next
	config.Previous = locationPage.Previous

	locationPage.Print()

	return nil
}
package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"time"
	"github.com/VirgiVaRu/pokedexcli/internal/pokecache"
	"github.com/VirgiVaRu/pokedexcli/internal/PokeAPI"
)

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Millisecond)
	config := &config{
				Next: "https://pokeapi.co/api/v2/location-area/",
				Previous: nil,
			}
	pokedex := &Pokedex{
		caughtPokemon: make(map[string]PokeAPI.Pokemon),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)

		if len(words) == 0 {
			continue
		}

		commandWritten := words[0]
		var parameters []string
		if len(words) > 1 {
			for i, word := range words {
				if i == 0 {
					continue
				}
				parameters = append(parameters, word)
			}
			
		}

		command, ok := getCommands()[commandWritten]
		if ok {
			err := command.callback(config, cache, parameters, pokedex)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}	
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	rawText := strings.Trim(lowerText, " ")
	split := strings.Split(rawText, " ")
	var clean []string
	for _, word := range split {
		if len(word) > 0 {
			clean = append(clean, word)
		}
	}
	return clean
}

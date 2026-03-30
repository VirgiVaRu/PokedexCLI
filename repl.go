package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)

		commandWritten := words[0]

		command, ok := getCommands()[commandWritten]
		if ok {
			err := command.callback()
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

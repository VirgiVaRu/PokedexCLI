package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	supportedCommands := registerCommands()

	for ;; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)

		command, ok := supportedCommands[words[0]]
		if !ok {
			fmt.Printf("Unknown command\n")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}
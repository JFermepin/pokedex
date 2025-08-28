package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	var configuration config = config{
		next:     "https://pokeapi.co/api/v2/location-area/",
		previous: "",
		page:     0,
	}

	var configPointer *config = &configuration

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(configPointer)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

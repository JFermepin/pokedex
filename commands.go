package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     string
	previous string
	page     int
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show help",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas",
			callback:    commandMapb,
		},
	}
	return commands
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	commands := getCommands()
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(config *config) error {

	data, err := getApiResponse(config.next)
	if err != nil {
		return err
	}

	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	config.next = data.Next
	config.page += 1
	if data.Previous != nil {
		config.previous = data.Previous.(string)
	}

	return nil
}

func commandMapb(config *config) error {
	if config.page == 1 {
		fmt.Println("You're on the first page")
		return nil
	} else if config.previous == "" {
		fmt.Println("No previous page available.")
		return nil
	} else {
		data, err := getApiResponse(config.previous)
		if err != nil {
			return err
		}

		for _, result := range data.Results {
			fmt.Println(result.Name)
		}

		return nil
	}

}

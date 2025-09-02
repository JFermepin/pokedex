package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	apistructs "github.com/Jfermepin/pokedex/internal/api-structs"
	"github.com/Jfermepin/pokedex/internal/pokecache"
)

func startRepl() {

	var cfg *config = &config{
		next:           "https://pokeapi.co/api/v2/location-area/",
		previous:       "",
		page:           0,
		locationsCache: pokecache.NewCache(5 * time.Second),
		pokemonsCache:  pokecache.NewCache(5 * time.Second),
		pokemonsCaught: map[string]apistructs.Pokemon{},
	}

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
			var err error

			if len(input) > 1 {
				err = command.callback(strings.Join(input[1:], " "), cfg)
			} else {
				err = command.callback("", cfg)
			}

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

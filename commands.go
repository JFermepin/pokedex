package main

import (
	"fmt"
	"math/rand/v2"
	"os"

	apistructs "github.com/Jfermepin/pokedex/internal/api-structs"
	"github.com/Jfermepin/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(string, *config) error
}

type config struct {
	next           string
	previous       string
	page           int
	locationsCache *pokecache.Cache
	pokemonsCache  *pokecache.Cache
	pokemonsCaught map[string]apistructs.Pokemon
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
		"explore": {
			name:        "explore",
			description: "Displays the names of the pokemons in the given location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch a pokemon. Takes the name of a Pokemon as an argument",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays the information of a caught pokemon. Takes the name of a caught Pokemon as an argument",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the names of the pokemons you have currently in your pokedex",
			callback:    commandPokedex,
		},
	}
	return commands
}

func commandExit(parameter string, config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(parameter string, config *config) error {
	commands := getCommands()
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(parameter string, config *config) error {

	data, err := getLocations(config)
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

func commandMapb(parameter string, config *config) error {
	if config.page == 1 {
		fmt.Println("You're on the first page")
		return nil
	} else if config.previous == "" {
		fmt.Println("No previous page available.")
		return nil
	} else {
		data, err := getLocations(config)
		if err != nil {
			return err
		}

		for _, result := range data.Results {
			fmt.Println(result.Name)
		}

		return nil
	}
}

func commandExplore(area string, config *config) error {
	fmt.Println("Exploring " + area + "...")
	data, err := getAreaPokemons(area, config)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemons:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Println("-" + encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(pokemon string, config *config) error {
	fmt.Println("Throwing a Pokeball at " + pokemon + "...")
	data, err := getPokemonInfo(pokemon, config)
	if err != nil {
		return err
	}

	chance := 1.0 / (1.0 + float64(data.BaseExperience)/50.0)

	if rand.Float64() < chance {
		fmt.Println(data.Name, "was caught!")
		config.pokemonsCaught[data.Name] = data
	} else {
		fmt.Println(data.Name, "escaped!")
	}

	return nil
}

func commandInspect(pokemon string, config *config) error {
	data, exists := config.pokemonsCaught[pokemon]
	if !exists {
		fmt.Println("You don't have that pokemon caught")
		return nil
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d", data.Name, data.Height, data.Weight)
	fmt.Println("Stats:")
	for _, s := range data.Stats {
		fmt.Printf("- %s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range data.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
	return nil
}

func commandPokedex(parameter string, config *config) error {
	if len(config.pokemonsCaught) == 0 {
		fmt.Println("You haven't caught any pokemons yet.")
		return nil
	}

	fmt.Println("Pokemons in your Pokedex:")
	for name := range config.pokemonsCaught {
		fmt.Println("- " + name)
	}
	return nil
}

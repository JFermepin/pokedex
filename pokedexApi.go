package main

import (
	"encoding/json"
	"io"
	"net/http"

	apistructs "github.com/Jfermepin/pokedex/internal/api-structs"
)

func getLocations(config *config) (apistructs.Locations, error) {
	var data apistructs.Locations

	if val, ok := config.locationsCache.Get(config.next); ok {
		err := json.Unmarshal(val, &data)
		if err != nil {
			return data, err
		}

		return data, nil
	}

	response, err := http.Get(config.next)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	config.locationsCache.Add(config.next, body)
	return data, nil
}

func getAreaPokemons(area string, config *config) (apistructs.LocationInfo, error) {
	var data apistructs.LocationInfo
	areaUrl := "https://pokeapi.co/api/v2/location-area/" + area

	if val, ok := config.pokemonsCache.Get(areaUrl); ok {
		err := json.Unmarshal(val, &data)
		if err != nil {
			return data, err
		}

		return data, nil
	}

	response, err := http.Get(areaUrl)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	config.pokemonsCache.Add(areaUrl, body)
	return data, nil
}

func getPokemonInfo(pokemon string, config *config) (apistructs.Pokemon, error) {
	var data apistructs.Pokemon
	pokemonUrl := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	if val, ok := config.pokemonsCache.Get(pokemonUrl); ok {
		err := json.Unmarshal(val, &data)
		if err != nil {
			return data, err
		}

		return data, nil
	}

	response, err := http.Get(pokemonUrl)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	config.pokemonsCache.Add(pokemonUrl, body)
	return data, nil
}

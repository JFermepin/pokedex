package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type PokeApiResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocations(config *config) (PokeApiResponse, error) {
	var data PokeApiResponse

	if val, ok := config.cache.Get(config.next); ok {
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

	config.cache.Add(config.next, body)
	return data, nil
}

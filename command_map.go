package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

type PokeApiJsonResponse struct{
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []struct{
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}


func commandMap() error{

	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil{
		fmt.Errorf("Error retrieving locations...", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil{
		return err
	}

	var PokemonLocations PokeApiJsonResponse
	if err := json.Unmarshal(data, &PokemonLocations); err != nil{
		fmt.Errorf("Error Unmarshalling data...")
	}
	results := PokemonLocations.Results[0]

	fmt.Println(results)

	return nil
}
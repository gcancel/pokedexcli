package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"io"
	"internal/pokecache"
)

func commandExplore(cfg *Config, cache pokecache.Cache, param string) error{
	
	exploreURL := "https://pokeapi.co/api/v2/location-area/" + param

	cacheData, exists := cache.Get(exploreURL)
	
	var PokemonLocationWithPokemon PokeApiLocationJsonResponse
	if exists{
		fmt.Printf("cached entry used...\n")
		for k,v := range cache.CacheEntries{
			fmt.Printf("cached entries: %v: %v\n", k, v.CreatedAt)
		}
	
		if err := json.Unmarshal(cacheData, &PokemonLocationWithPokemon); err != nil {
			fmt.Errorf("Error Unmarshalling data...")
		}
	}else{
		res, err := http.Get(exploreURL)
		if err != nil {
			fmt.Errorf("Error retrieving locations... %v", err)
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		// adding data to cache
		cache.Add(exploreURL, data)

	    // unmarshalling data
		if err := json.Unmarshal(data, &PokemonLocationWithPokemon); err != nil {
			fmt.Errorf("Error Unmarshalling data...")
		}
	}

	pokemonList := PokemonLocationWithPokemon.PokemonEncounters

	for _,pokemon := range pokemonList{
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
package main

import (
	"encoding/json"
	"fmt"
	"internal/pokecache"
	"io"
	"net/http"
)

func commandMap(cfg *Config, cache *pokecache.Cache) error {

	// check cache for data
	// if data is there jump to unmarshalling the JSON
	// 	--update the config
	// 	--print the results
	// if data is not there create a response with http.Get()
	// 	-- read the data with io.ReadAll()
	//  -- cache new data
	// 	-- begin unmarshalling data
	//  -- update config
	//  -- print results


	if cfg.Next == "" {
		fmt.Println("This is the last result page...")
		return nil
	}

	cacheData, exists := cache.Get(cfg.Next)
	var PokemonLocations PokeApiJsonResponse
	if exists{
		fmt.Printf("cached entry used...\n")
		for k,v := range cache.CacheEntries{
			fmt.Printf("cached entries: %v: %v\n", k, v.CreatedAt)
		}
	
		if err := json.Unmarshal(cacheData, &PokemonLocations); err != nil {
			fmt.Errorf("Error Unmarshalling data...")
		}
	}else{
		res, err := http.Get(cfg.Next)
		if err != nil {
			fmt.Errorf("Error retrieving locations... %v", err)
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		// adding data to cache
		cache.Add(cfg.Next, data)

	    // unmarshalling data
		if err := json.Unmarshal(data, &PokemonLocations); err != nil {
			fmt.Errorf("Error Unmarshalling data...")
		}
	}
	
	// updating the config variable to point to the next page
	cfg.Prev = PokemonLocations.Previous
	cfg.Next = PokemonLocations.Next
	results := PokemonLocations.Results

	for _, loc := range results {
		fmt.Printf("%s\n", loc.Name)
	}
	fmt.Println(PokemonLocations.Next)

	return nil
}

func commandMapB(cfg *Config, cache *pokecache.Cache) error {

	if cfg.Prev == "" {
		fmt.Println("This is the first result page...")
		return nil
	}

	cacheData, exists := cache.Get(cfg.Prev)
	var PokemonLocations PokeApiJsonResponse
	if exists{
		fmt.Printf("cached entry used...\n")
		for k,v := range cache.CacheEntries{
			fmt.Printf("cached entries: %v: %v\n", k, v.CreatedAt)
		}

		if err := json.Unmarshal(cacheData, &PokemonLocations); err != nil {
			fmt.Errorf("Error Unmarshalling data...")
		}
	}else{
		res, err := http.Get(cfg.Prev)
		if err != nil {
			fmt.Errorf("Error retrieving locations... %v", err)
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		// adding data to cache
		cache.Add(cfg.Prev, data)

	    // unmarshalling data
		if err := json.Unmarshal(data, &PokemonLocations); err != nil {
			fmt.Errorf("Error Unmarshalling data...")
		}
	}
	
	// updating the config variable to point to the next page
	cfg.Prev = PokemonLocations.Previous
	cfg.Next = PokemonLocations.Next
	results := PokemonLocations.Results
    
	// printing page results
	for _, loc := range results {
		fmt.Printf("%s\n", loc.Name)
	}

	return nil
}

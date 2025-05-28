package main

import(
	"fmt"
	"io"
	"encoding/json"
	"net/http"
	"time"
	"math/rand"
	"internal/pokecache"
)

func commandCatch(cfg *Config, cache pokecache.Cache, param string) error{
	// base functionality is there, just need to add caching and catch randomization...
	url := "https://pokeapi.co/api/v2/pokemon/" + param
	var pokemon Pokemon
	

	cachedData, exists := cache.Get(url)
	if exists{
		fmt.Printf("using cached data: %v", url)
		if err := json.Unmarshal(cachedData, &pokemon); err != nil{
			// error in unmarshalling data, most likely invalid pokemon
			return fmt.Errorf("Pokemon not found in Pokedex... %v", err)
		}
	}else{
		res, err := http.Get(url)
		if err != nil{
			return fmt.Errorf("Error retrieving specified pokemon...")
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil{
			return err
		}
		cache.Add(url, data)

		if err := json.Unmarshal(data, &pokemon); err != nil{
			// error in unmarshalling data, most likely invalid pokemon
			return fmt.Errorf("Pokemon not found in Pokedex... %v", err)
		}
	}

	// creating catch randomizer based on base experience of pokemon

	fmt.Printf("Throwing a pokeBag at %v...\n", pokemon.Name)
	time.Sleep(3 * time.Second)

	catch := rand.Intn(pokemon.BaseExperience)
	if catch > int(pokemon.BaseExperience / 2){
		fmt.Printf("%v was caught!\n", pokemon.Name)
		// adding new pokemon to backpack
		// this needs to be a pointer to the bag and moved out of the function
		cfg.Pokedex[pokemon.Name] = pokemon

	 	for _,pokeball := range cfg.Pokedex{
			fmt.Printf("%v\n", pokeball.Name)
		}
	}else{
		// pokemon escaped
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	return nil
}


package main

import(
	"fmt"
	"internal/pokecache"
)

func commandPokedex(cfg *Config, cache pokecache.Cache, param string) error{
	caughtPokemon := cfg.Pokedex
	if len(caughtPokemon) >= 1{
		for _,pokemon := range caughtPokemon{
			fmt.Printf(" -%v\n", pokemon.Name)
		}
	}else{
		fmt.Println("you have no captured pokemon")
	}
	return nil
}
package main

import(
	"fmt"
	"internal/pokecache"
)

func commandInspect(cfg *Config, cache pokecache.Cache, param string) error{
	pokemon := param
	caughtPokemon, exists := cfg.Pokedex[pokemon]
	if exists{
		fmt.Printf("Name: %v\n", caughtPokemon.Name)
		fmt.Printf("Height: %v\n", caughtPokemon.Height)
		fmt.Printf("Weight: %v\n", caughtPokemon.Weight)
		fmt.Println("Stats:")
		pokeStats := caughtPokemon.Stats
		for _,stat := range pokeStats{
			fmt.Printf(" -%v: %v\n",stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		pokeTypes := caughtPokemon.Types
		for _,t := range pokeTypes{
			fmt.Printf(" - %v\n", t.Type.Name)
		}
	}else{
		fmt.Printf("you have not caught that pokemon\n")
	}
	return nil
}
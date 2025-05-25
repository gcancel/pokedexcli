package main

import(
	"fmt"
	"internal/pokecache"
)

func commandHelp(cfg *Config, c *pokecache.Cache) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	printCommands()
	return nil
 }
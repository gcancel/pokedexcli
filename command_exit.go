package main

import(
	"fmt"
	"os"
	"internal/pokecache"
)

func commandExit(cfg *Config, c pokecache.Cache, param string) error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
 }
package main

import(
	"fmt"
)

func commandHelp(cfg *Config) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	printCommands()
	return nil
 }
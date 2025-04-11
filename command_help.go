package main

import(
	"fmt"
)

func commandHelp() error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	printCommands()
	return nil
 }
package main

 import(
	"fmt"
	"strings"
	"bufio"
	"os"
	"time"
	"internal/pokecache"
 )

 func startRepl(){
	scanner := bufio.NewScanner(os.Stdin)

	// creating cache
	cache := pokecache.NewCache(10 * time.Second)
	
	for{
		fmt.Print("Pokedex > ")
		scanner.Scan()

		// splitting the input and capturing the command (first string in the slice)
		userInput := cleanInput(scanner.Text())
		userCommand := strings.TrimSpace(userInput[0])

		var userParameter string
		if len(userInput) == 0{
			continue
		}
		// if there is a parameter passed, pass it down to the commands
		if len(userInput) > 1{
			userParameter = strings.TrimSpace(userInput[1])
		}
		
		// execute the user command
		executeCommands(userCommand, userParameter, cache)

		// fmt.Printf("Your command was: %s \n", userCommand)
	}
 }

 func cleanInput(text string) []string{
	stringSlice := strings.Split(strings.ToLower(text), " ")
	return stringSlice
 }

 type cliCommand struct{
	name string
	description string
	callback func(cfg *Config, c pokecache.Cache, param string) error
 }

 func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"exit":{
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help":{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map":{
			name: "map",
			description: "Displays locations in the Pokemon World - Forward",
			callback: commandMap,
		},
		"mapb":{
			name: "mapb",
			description: "Displays locations in the Pokemon World - Back",
			callback: commandMapB,
		},
		"explore":{
			name: "explore",
			description: "Explores specified location.",
			callback: commandExplore,
		},
	 }
 }

 func printCommands(){
	for _,cmd := range getCommands(){
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
 }

 func executeCommands(cmd, param string, c pokecache.Cache){

	execute, exists := getCommands()[cmd];
	if exists{
		err := execute.callback(&APIConfig, c, param)
		if err != nil{
			fmt.Println(err)
			return
		}
		return
	}
	fmt.Println("Unknown Command")
	return
 }
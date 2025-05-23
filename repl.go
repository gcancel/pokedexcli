package main

 import(
	"fmt"
	"strings"
	"bufio"
	"os"
 )

 func startRepl(){
	scanner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		//splitting the input and capturing the command (first string in the slice)
		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0{
			continue
		}
		userCommand := strings.TrimSpace(userInput[0])
		//fmt.Println(userCommand)
		//execute the user command
		executeCommands(userCommand)

		//fmt.Printf("Your command was: %s \n", userCommand)
	}
 }

 func cleanInput(text string) []string{
	stringSlice := strings.Split(strings.ToLower(text), " ")
	return stringSlice
 }

 type cliCommand struct{
	name string
	description string
	callback func(cfg *Config) error
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
			description: "Displays locations in the Pokemon World",
			callback: commandMap,
		},
		"mapb":{
			name: "mapb",
			description: "Displays locations in the Pokemon World",
			callback: commandMapB,
		},
	 }
 }

 func printCommands(){
	for _,cmd := range getCommands(){
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
 }

 func executeCommands(cmd string){

	execute, exists := getCommands()[cmd];
	if exists == false{
		fmt.Println("Unknown Command")
		return
	}
	err := execute.callback(&APIConfig)
	if err != nil{
		fmt.Println(err)
		return
	}
 }
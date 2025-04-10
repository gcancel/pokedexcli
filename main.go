package main
 import(
	"fmt"
	"strings"
	"bufio"
	"os"
 )

 //commands
 type cliCommand struct{
	name string
	description string
	callback func()
 }
 var commands = map[string]cliCommand{
	"exit":{
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help":{
		name: "help",
		description: "Displays the commands and their usage for the Pokedex",
		callback: commandHelp,
	},
 }

 //functions

 func cleanInput(text string) []string{
	stringSlice := strings.Split(strings.ToLower(text), " ")
	return stringSlice
 }

 func commandExit(){
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
 }

 func commandHelp(){
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for command,_ := range commands{
		fmt.Printf("%s: %s \n", command, commands[command].description)
	}
 }

 func executeCommands(cmd string){
	execute, exists := commands[cmd];
	if !exists{
		fmt.Println("Unknown Command")
	}
	execute.callback()
 }

 func main(){
	
	for{
		fmt.Print("Pokedex > ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println(err, "error occurred during input")
		}
		//splitting the input and capturing the command (first string in the slice)
		userInput := cleanInput(input)
		userCommand := userInput[0]

		//execute the user command
		executeCommands(userCommand)

		//fmt.Printf("Your command was: %s \n", userCommand)
	}
 }
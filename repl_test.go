package main

import(
	"testing"
)

func TestCleanInput(t *testing.T){

	cases := []struct{
		input string
		expected []string
	}{
		{
			input: "hello world", 
			expected: []string{"hello", "world"},
		},
		{
			input: "Hello, World!",
			expected: []string{"hello,", "world!"},
		},
		{
			input: "Greg Martin",
			expected: []string{"greg", "martin"},
		},
		{
			input: "this is a sentence",
			expected: []string{"this", "is", "a", "sentence"},
		},
		// add more cases here
	}
	
	for _,c := range cases{
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("length of actual slice: %v does not equal expected slice: %v ", actual, c.expected)
		}

		for i:= range actual{
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord{
				t.Errorf("word: %s does not equal expectedWord: %s ", word, expectedWord)
			}
		}
		
	}

}

func TestPrintCommands(t *testing.T){
	cases := []struct{
		input  func() map[string]cliCommand
		expected map[string]cliCommand
	}{
		{
			input: getCommands,
			expected: map[string]cliCommand{
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
			 },
		},
		
	}
	for _,c := range cases{
		if c.input()["exit"].name != c.expected["exit"].name{
			t.Errorf("Map does not match existing map: %s, %s", c.input()["exit"].name, c.expected["exit"].name)
		}
	}
	
}



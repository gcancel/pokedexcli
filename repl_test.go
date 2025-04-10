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


package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

func commandMapB(cfg *Config) error{

	_, offset, _ := ParsePageLimit(cfg.Next)
	if offset - 20 == 0 {
		fmt.Println("This is the first result page...")
		return nil
	}

	res, err := http.Get(cfg.Prev)
	if err != nil{
		fmt.Errorf("Error retrieving locations...", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil{
		return err
	}

	var PokemonLocations PokeApiJsonResponse
	if err := json.Unmarshal(data, &PokemonLocations); err != nil{
		fmt.Errorf("Error Unmarshalling data...")
	}
	//fmt.Println(PokemonLocations.Count)

	//need to add guard clause to check if the offset is equal to zero, if so this is the first page
	//updating the config variable to point to the next page
	cfg.Prev = PokemonLocations.Previous
	cfg.Next = PokemonLocations.Next
	results := PokemonLocations.Results

	for _,loc := range results{
		fmt.Printf("%s\n", loc.Name)
	}
	//fmt.Println(PokemonLocations.Next)

	return nil
}
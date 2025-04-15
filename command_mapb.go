package main

import(
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

func commandMapB(cfg *Config) error{

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
	fmt.Println(PokemonLocations.Count)

	//updating the config variable to point to the next page
	cfg.Prev = PokemonLocations.Previous
	cfg.Next = PokemonLocations.Next
	results := PokemonLocations.Results

	for _,loc := range results{
		fmt.Printf("%s\n", loc.Name)
	}
	fmt.Println(PokemonLocations.Next)

	return nil
}
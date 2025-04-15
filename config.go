package main


type PokeApiJsonResponse struct{
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []struct{
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}

type Config struct{
	Next string
	Prev string
}

var APIConfig = Config{
	Next: "https://pokeapi.co/api/v2/location-area",
	Prev: "",
}
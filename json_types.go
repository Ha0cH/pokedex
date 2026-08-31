package main

// Paginated location areas
type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreasPage struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

// LocationArea endpoint detail
type PokemonNameURL struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon PokemonNameURL `json:"pokemon"`
}

type LocationAreaDetail struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

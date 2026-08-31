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

// Pokemon endpoint
type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int             `json:"base_stat"`
	Stat     PokemonStatInfo `json:"stat"`
}

type PokemonStatInfo struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Type PokemonTypeInfo `json:"type"`
}

type PokemonTypeInfo struct {
	Name string `json:"name"`
}

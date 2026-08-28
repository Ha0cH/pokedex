package main

type locationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type locationAreasPage struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}

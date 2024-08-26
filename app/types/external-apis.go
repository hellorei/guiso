package types

type PokemonList struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Pokemon `json:"results"`
}

type PokemonSprites struct {
	FrontDefault string `json:"front_default"`
}

type Pokemon struct {
	Name    string         `json:"name"`
	Sprites PokemonSprites `json:"sprites"`
}

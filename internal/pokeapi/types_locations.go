package pokeapi

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationAreaResult struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}

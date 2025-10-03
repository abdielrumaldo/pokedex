package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (PokemonResult, error) {

	url := baseUrl + "/pokemon/" + pokemon
	pokemonResp := PokemonResult{}

	cachedResp, ok := c.httpCache.Get(url)
	if ok != false {
		err := json.Unmarshal(cachedResp, &pokemonResp)
		if err != nil {
			return pokemonResp, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemonResp, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokemonResp, err
	}
	if resp.StatusCode != 200 {
		return pokemonResp, fmt.Errorf("Pokemon '%s': not found", pokemon)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemonResp, err
	}

	// add to cache
	c.httpCache.Add(url, data)

	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return pokemonResp, err
	}

	return pokemonResp, err
}

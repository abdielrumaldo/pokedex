package pokeapi

import (
	"net/http"
	"time"

	"github.com/abdielrumaldo/pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	httpCache  pokecache.Cache
	Pokedex    map[string]PokemonResult
}

func NewClient(timeout time.Duration) Client {
	cache := pokecache.NewCache(5 * time.Minute)
	pokedex := make(map[string]PokemonResult)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		httpCache: cache,
		Pokedex:   pokedex,
	}
}

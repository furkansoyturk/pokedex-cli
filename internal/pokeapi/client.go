package pokeapi

import (
	"github.com/furkansoyturk/pokedex-cli/internal/pokecache"
	"net/http"
	"time"
)

type Client struct {
	cache      pokecache.CacheStruct
	httpClient http.Client
}

func NewClient(timeout, timeToLive time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(timeToLive),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

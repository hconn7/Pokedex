package pokeapi

import (
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
}

// NewClient -
func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

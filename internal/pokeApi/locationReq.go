package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreaResp{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}

	if resp.StatusCode > 299 {
		return LocationAreaResp{}, fmt.Errorf("error, status code: %v", resp.StatusCode)
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}
	locationResp := LocationAreaResp{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return locationResp, nil
}

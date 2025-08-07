package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (locationAreaResult, error) {

	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreaResult{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaResult{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaResult{}, err
	}

	locationResp := locationAreaResult{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return locationAreaResult{}, err
	}

	return locationResp, nil
}

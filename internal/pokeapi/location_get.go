package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(area string) (LocationResult, error) {

	url := baseUrl + "/location-area/" + area
	locationResp := LocationResult{}

	cachedResp, ok := c.httpCache.Get(url)
	if ok != false {
		err := json.Unmarshal(cachedResp, &locationResp)
		if err != nil {
			return LocationResult{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationResp, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationResp, err
	}
	if resp.StatusCode != 200 {
		return locationResp, fmt.Errorf("Area '%s': not found", area)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationResp, err
	}

	// add to cache
	c.httpCache.Add(url, data)

	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return locationResp, err
	}

	return locationResp, err
}

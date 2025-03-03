package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetLocationListData(url string) (LocationListResponseFormat, error) {
	
	req, req_err := http.NewRequest("GET", url, nil)
	if req_err != nil {
		return LocationListResponseFormat{}, req_err
	}

	res, res_err := c.httpClient.Do(req)
	if res_err != nil {
		return LocationListResponseFormat{}, res_err
	}

	defer res.Body.Close()
	var output LocationListResponseFormat
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&output); err != nil {
		return LocationListResponseFormat{}, err
	}
	
	return output, nil
}

func (c *Client) GetLocationSpecificData(locationName string) (LocationSpecificResponseFormat, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + locationName
	
	req, req_err := http.NewRequest("GET", url, nil)
	if req_err != nil {
		return LocationSpecificResponseFormat{}, req_err
	}

	res, res_err := c.httpClient.Do(req)
	if res_err != nil {
		return LocationSpecificResponseFormat{}, res_err
	}

	defer res.Body.Close()
	var output LocationSpecificResponseFormat
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&output); err != nil {
		return LocationSpecificResponseFormat{}, nil
	}

	return output, nil
}

func (c *Client) GetPokemonResponseData (pokemonName string) (PokemonResponseFormat, error) {

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	req, req_err := http.NewRequest("GET", url, nil)
	if req_err != nil {
		return PokemonResponseFormat{}, req_err
	}

	res, res_err := c.httpClient.Do(req)
	if res_err != nil {
		return PokemonResponseFormat{}, res_err
	}

	defer res.Body.Close()
	var output PokemonResponseFormat
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&output); err != nil {
		return PokemonResponseFormat{}, nil
	}

	return output, nil
}
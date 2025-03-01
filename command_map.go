package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"errors"
)

func commandMap(con *config) error {
	
	mapData, err := getData(con.Next)
	
	if err != nil {
		return err
	}
	for _, item := range mapData.Results {
		fmt.Println(item.Name)
	}

	con.Next = mapData.Next
	con.Previous = mapData.Previous
	
	return nil
}

func commandMapb(con *config) error {
	
	if len(con.Previous) == 0 {
		return errors.New("There is no previous page to view")
	}

	mapData, err := getData(con.Previous)
	
	if err != nil {
		return err
	}
	for _, item := range mapData.Results {
		fmt.Println(item.Name)
	}

	con.Next = mapData.Next
	con.Previous = mapData.Previous
	
	return nil
}


func getData(url string) (ResponseFormat, error) {
	
	req, req_err := http.NewRequest("GET", url, nil)
	if req_err != nil {
		return ResponseFormat{}, req_err
	}

	client := http.Client{}
	res, res_err := client.Do(req)
	if res_err != nil {
		return ResponseFormat{}, res_err
	}

	defer res.Body.Close()
	var output ResponseFormat
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&output); err != nil {
		return ResponseFormat{}, err
	}
	
	return output, nil
}

type NameUrlPair struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type ResponseFormat struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []NameUrlPair `json:"results"`
}









func getStringData(url string) (string, error) {
	
	req, req_err := http.NewRequest("GET", url, nil)
	if req_err != nil {
		return "", req_err
	}

	client := http.Client{}
	res, res_err := client.Do(req)
	if res_err != nil {
		return "", res_err
	}

	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
    		bodyBytes, err := io.ReadAll(res.Body)
	    if err != nil {
	       	return "", err
       		}
	    bodyString := string(bodyBytes)
	    return bodyString, nil
	}
	return "", nil
}

type LocationArea struct {
	Id int `json:"id"`
	Name string `json:"name"`
	GameIndex int `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				Url string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location map[string]string `json:"location"`
	Names []struct {
		Name string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"language"`
	}
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				Url string `json:"url"`
			} `json:"version"`
			MaxChance string `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				ConditionValues []struct {
					Name string `json:"name"`
					Url string `json:"url"`
				} `json:"condition_values"`
				Chance int `json:"chance"`
				Method struct {
					Name string `json:"name"`
					Url string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

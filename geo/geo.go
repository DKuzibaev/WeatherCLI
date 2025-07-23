package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponse struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	
	if city != "" {
		if isCity := checkCity(city); isCity {
			panic("There is no such city!")
		}
		return &GeoData{
			City: city,
		}, nil
	}
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	locUrl, err := CheckEnv("CITY_LOC")

	if err != nil {
		return nil, err
	}

	resp, err := http.Get(locUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("NOT-200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

func CheckEnv(path string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", err
	}

	envVar := os.Getenv(path)
	return envVar, nil
}


func checkCity(city string) (bool) {
	cityUrl, err := CheckEnv("CITY_POPLR")

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	postBody, _ :=json.Marshal(map[string]string{
		"city": city,
	})

	resp, err := http.Post(cityUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var populationResponse CityPopulationResponse
	json.Unmarshal(body, &populationResponse)
	return populationResponse.Error

}
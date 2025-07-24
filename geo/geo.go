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

var (
	ErrNoCity       = errors.New("NOCITY")
	ErrStatusNot200 = errors.New("NOT-200")
	ErrCityPop      = errors.New("CITY_POPLR IS EMPTY")
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponce struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {

	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			return nil, ErrNoCity
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
		return nil, ErrStatusNot200
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

func checkCity(city string) bool {
	cityUrl, err := CheckEnv("CITY_POPLR")
	if err != nil {
		return false
	}

	postBody, _ := json.Marshal(map[string]string{
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
	var populationResponce CityPopulationResponce
	json.Unmarshal(body, &populationResponce)
	return !populationResponce.Error

}

func CheckEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
	return val, nil
}

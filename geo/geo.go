package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {

	if len(city) != 0 {
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

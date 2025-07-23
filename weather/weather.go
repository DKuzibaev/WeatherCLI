package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"weather-cli/geo"

	"github.com/joho/godotenv"
)

func GetWeather(geo geo.GeoData, format int) (string, error) {
	wttrUrl, err := CheckEnv("WTTR_ADRT")
	if err != nil {
		return "", err
	}

	baseUrl, err := url.Parse(wttrUrl + geo.City)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	params := url.Values{}
	params.Add("format", fmt.Sprintf("%d", format))

	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return string(body), nil
}

func CheckEnv(path string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", err
	}
	envVar := os.Getenv(path)
	return envVar, nil
}

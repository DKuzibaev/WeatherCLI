package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"weather-cli/geo"
)

var (
	ErrWrongFormat = errors.New("WRONG_FORMAT")
	ErrInvalidUrl  = errors.New("ERROR_URL")
	ErrHttpRequest = errors.New("ERROR_HTTP")
	ErrReadBody    = errors.New("ERROR_READ_BODY")
)

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrWrongFormat
	}

	wttrUrl, err := CheckEnv("WTTR_ADRT")
	if err != nil {
		return "", err
	}

	baseUrl, err := url.Parse(wttrUrl + geo.City)

	if err != nil {
		fmt.Println(err.Error())
		return "", ErrInvalidUrl
	}

	params := url.Values{}
	params.Add("format", fmt.Sprintf("%d", format))

	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())

	if err != nil {
		fmt.Println(err.Error())
		return "", ErrHttpRequest
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", ErrReadBody
	}

	return string(body), nil
}

func CheckEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
	return val, nil
}

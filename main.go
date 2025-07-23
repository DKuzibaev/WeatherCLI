package main

import (
	"flag"
	"fmt"
	"log"
	"weather-cli/geo"
	"weather-cli/weather"
)

func main() {
	log.Println("Starting WeatherCLI...")

	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(geoData)

	weatherData, err := weather.GetWeather(*geoData, *format)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(weatherData)
}

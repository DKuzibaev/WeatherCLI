package main

import (
	"flag"
	"fmt"
	"log"
	"weather-cli/geo"
)

func main() {
	log.Println("Starting WeatherCLI...")

	city := flag.String("city", "", "Город пользователя")
	//format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)

	deoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(deoData)
	//fmt.Println(*format)
}

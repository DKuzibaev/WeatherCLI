package main

import (
	"flag"
	"fmt"
	"log"
	"weather-cli/geo"
	"weather-cli/weather"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

var menuValues = []string{
	"1. Запустить определение по IP",
	"2. Указать город и формат (от 1 до 4)",
	"3. Выход",
}

func printMenu(menu ...string) {
	for _, item := range menu {
		fmt.Println(item)
	}
}

func runMenu() {
	for {
		printMenu(menuValues...)
		fmt.Print("Выбрать пункт: ")
		var point string
		fmt.Scanln(&point)

		switch point {
		case "1":
			startWithNoFlags()
		case "2":
			startWithCityFags()
		case "3":
			fmt.Println("Выход из программы")
			return // Выход из функции и программы
		default:
			fmt.Println("Неверный пункт меню, попробуйте снова")
		}
	}
}

func main() {
	fmt.Println("-----Терминальная погода-----")
	fmt.Println("")
	runMenu()
}

func startWithNoFlags() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 4, "Формат вывода погоды")
	flag.Parse()
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	weatherData, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(weatherData)
}

func startWithCityFags() {
	var scanCity string
	var scanFormat int

	fmt.Print("Введите город: ") // Важно: Print, а не Println!
	_, err := fmt.Scanln(&scanCity)
	if err != nil {
		fmt.Println("Ошибка ввода города:", err)
		return
	}

	fmt.Print("Введите формат: ") // Аналогично
	_, err = fmt.Scanln(&scanFormat)
	if err != nil {
		fmt.Println("Ошибка ввода формата:", err)
		return
	}

	city := flag.String("city", scanCity, "Город пользователя")
	format := flag.Int("format", scanFormat, "Формат вывода погоды")
	flag.Parse()
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	weatherData, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(weatherData)
}

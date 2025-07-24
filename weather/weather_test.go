package weather_test

import (
	"os"
	"strings"
	"testing"
	"weather-cli/geo"
	"weather-cli/weather"
)

func TestGetWeather(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функиции
	_ = os.Setenv("WTTR_ADRT", "https://wttr.in/")
	expected := "Moscow"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3
	// Act - выполняем функцию
	result, err := weather.GetWeather(geoData, format)

	// Assert - проверка результата с  expected
	if err != nil {
		t.Errorf("Пришла ошибка %v", err)
	}

	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, Получено - %v", expected, result)
	}

}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 147},
	{name: "Zero format", format: 0},
	{name: "Negative format", format: -5},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	_ = os.Setenv("WTTR_ADRT", "https://wttr.in/")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "Moscow"
			geoData := geo.GeoData{
				City: expected,
			}
			_, err := weather.GetWeather(geoData, tc.format)
			if err != weather.ErrWrongFormat {
				t.Errorf("Ожидалось %v, Получено - %v", weather.ErrWrongFormat, err)
			}
		})
	}
}

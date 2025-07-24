package geo_test

import (
	"os"
	"testing"
	"weather-cli/geo"
)

func TestGetMyLocation(t *testing.T) {

	// Arrange - подготовка, expected результат, данные для функиции
	_ = os.Setenv("CITY_POPLR", "https://countriesnow.space/api/v0.1/countries/population/cities")
	city := "Moscow"
	expected := geo.GeoData{
		City: "Moscow",
	}
	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)

	// Assert - проверка результата с  expected
	if err != nil {
		t.Errorf("Ошибка получаения города: %s", err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, Получено - %v", expected, got)
	}

}

func TestGetMyLocationNoCity(t *testing.T) {
	// arrange - подготовка, expected результат, данные для функиции
	city := "Lsjsdaaada"
	// act - выполняем функцию
	_, err := geo.GetMyLocation(city)

	// assert - проверка результата с  expected
	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v, получено %v", geo.ErrNoCity, err)
	}
}

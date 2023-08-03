package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	loggers "github.com/quangduoong/WeatherCLIApp/src/helpers"
	"github.com/quangduoong/WeatherCLIApp/src/models"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		loggers.Fatal(err)
	}
}

func main() {
	apiKey := os.Getenv("API_KEY")
	q := "auto:ip"

	if len(os.Args) > 1{
		q = os.Args[1]
	}

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + apiKey + "&q=" + q + "&days=1&aqi=no&alerts=no")
	loggers.PanicIfNotNil(err)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		loggers.Log("Weather API not available.")
		return
	}

	body, err := io.ReadAll(res.Body)
	loggers.PanicIfNotNil(err)

	var weather models.Weather
	err = json.Unmarshal(body, &weather)
	loggers.PanicIfNotNil(err)

	location, current := weather.Location, weather.Current

	fmt.Printf(
		"%s, %s: %.0fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text)
}

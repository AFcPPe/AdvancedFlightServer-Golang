package util

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

var WeatherData sync.Map

func SyncWeather() {
	log.Printf("Getting weather data from %s", viper.GetString("server.weather_url"))
	resp, err := http.Get(viper.GetString("server.weather_url"))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	for _, v := range strings.Split(string(body), "\n") {
		if len(v) > 4 {
			WeatherData.Store(v[:4], v)

		}
	}
	log.Println("Weather data updated")
}

func GetWeatherByICAO(ICAO string) string {
	value, has := WeatherData.Load(ICAO)
	if has {
		return value.(string)
	}
	return ""
}

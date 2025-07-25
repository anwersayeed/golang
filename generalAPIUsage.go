package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const apiKey = "y847r843r73489r7394r7"

func fetchWeather(city string) interface{} {
	var data struct {
		Main struct {
			Temp float64 json:"temp"
		} json:"main"
	}

	url := fmt.Sprintf("cdnjkfdjncdjfkncvkf/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding for %s: %s\n", city, err)
		return data
	}

	return data
}

func main() {

	startNow := time.Now()

	cities := []string{"Toronot", "London"}
	for _, city := range cities {
		data := fetchWeather(city)
		fmt.Println("This is the data", data)
	}

	fmt.Println("This operation took:", time.Since(startNow))
}
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const apiKey = "y847r843r73489r7394r7"

func fetchWeather(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {
	var data struct {
		Main struct {
			Temp float64 json:"temp"
		} json:"main"
	}

	defer wg.Done() // ensures

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

	ch <- fmt.Sprintf("This is the %s", city)

	return data
}

func main() {

	startNow := time.Now()

	cities := []string{"Toronot", "London"}

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchWeather(city, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("This operation took:", time.Since(startNow))
}
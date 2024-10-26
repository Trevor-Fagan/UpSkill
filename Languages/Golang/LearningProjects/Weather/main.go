package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Location: ")
		location, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		location = strings.TrimSpace(location)

		queryString := "http://api.weatherstack.com/current?access_key=" + os.Getenv("API_KEY") + "&query=" + location + "&units=f"
		weatherResp := getWeather(queryString)
		outputWeather(weatherResp)
	}
}

func getWeather(queryString string) []byte {
	resp, err := http.Get(queryString)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return bodyBytes
}

func outputWeather(weatherResp []byte) {
	var result map[string]interface{}
	_ = json.Unmarshal([]byte(weatherResp), &result)

	if resultLocation, ok := result["location"].(map[string]interface{}); ok {
		fmt.Println(resultLocation["name"])
	}

	if current, ok := result["current"].(map[string]interface{}); ok {
		fmt.Printf("Temperature: %f F\n", current["temperature"])
	}
}

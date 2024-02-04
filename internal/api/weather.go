package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetWeatherInfo(apiKey, location string) (string, error) {
	apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", location, apiKey)

	response, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	// Unmarshalling
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return "", err
	}

	var currTemp float64
	var description string
	// Getting current temp
	if main, ok := result["main"].(map[string]interface{}); ok {
		if temp, ok := main["temp"].(float64); ok {
			currTemp = temp - 273.15
		}
	}
	// Getting weather description
	if weather, ok := result["weather"].([]interface{}); ok {
		if len(weather) > 0 {
			if description, ok = weather[0].(map[string]interface{})["description"].(string); ok {
			}
		}
	}

	info := fmt.Sprintf("Сurrent weather in %s: %s, temperature %d °C", strings.Title(location), description, int(currTemp))

	return info, nil
}

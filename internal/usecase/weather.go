package usecase

import (
	"discord-bot/internal/api"
)

type WeatherUsecase struct {
	apiKey string
}

// Creates new WeatherUsecase
func NewWeatherUsecase(apiKey string) *WeatherUsecase {
	return &WeatherUsecase{
		apiKey: apiKey,
	}
}

// Getting weather information
func (w *WeatherUsecase) GetWeather(location string) (string, error) {
	info, err := api.GetWeatherInfo(w.apiKey, location)
	if err != nil {
		return "", err

	}
	return info, nil
}

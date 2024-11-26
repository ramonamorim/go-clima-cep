package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWeather(t *testing.T) {
	city := "Joinville"
	apiKey := "0c40606abeb649ffb5903826242611"

	weatherClient := NewWeatherClient(apiKey)
	weather, err := weatherClient.GetWeather(city)
	assert.NoError(t, err)

	assert.NotEmpty(t, weather.Condition)
	assert.NotZero(t, weather.Celsius)

	t.Logf("Weather in %s: %s, %.2f°C\n", city, weather.Condition, weather.Celsius)
}

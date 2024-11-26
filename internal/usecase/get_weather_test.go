package usecase

import (
	"testing"

	"github.com/ramonamorim/go-clima-cep/internal/infra/client"
	"github.com/stretchr/testify/assert"
)

func TestGetWeatherUseCase_Execute_Success(t *testing.T) {
	ApiKey := "0c40606abeb649ffb5903826242611"
	cepClient := client.NewCepClient()
	weatherClient := client.NewWeatherClient(ApiKey)

	useCase := NewGetWeatherUseCase(cepClient, weatherClient)

	input := GetWeatherRequest{CEP: "89221220"}
	output, err := useCase.Execute(input)

	assert.NoError(t, err)
	assert.NotEmpty(t, output.Celsius)
	assert.NotEmpty(t, output.Kelvin)
	assert.NotEmpty(t, output.Fahrenheit)
}

func TestGetWeatherUseCase_Execute_InvalidCEPError(t *testing.T) {
	ApiKey := "0c40606abeb649ffb5903826242611"
	cepClient := client.NewCepClient()
	weatherClient := client.NewWeatherClient(ApiKey)

	useCase := NewGetWeatherUseCase(cepClient, weatherClient)

	input := GetWeatherRequest{CEP: "99999999"}
	output, err := useCase.Execute(input)

	assert.Error(t, err)
	assert.Empty(t, output)
}

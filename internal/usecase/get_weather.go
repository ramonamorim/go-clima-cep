package usecase

import (
	"errors"

	"github.com/ramonamorim/go-clima-cep/internal/entity"
	"github.com/ramonamorim/go-clima-cep/internal/infra/client"
)

type GetWeatherRequest struct {
	CEP string `json:"cep"`
}

type GetWeatherResponse struct {
	Celsius    string `json:"temp_C"`
	Fahrenheit string `json:"temp_F"`
	Kelvin     string `json:"temp_K"`
}

type GetWeatherUseCase struct {
	CEPClient     client.CepClientInterface
	WeatherClient client.WeatherClientInterface
}

func NewGetWeatherUseCase(
	CEPClient client.CepClientInterface,
	WeatherClient client.WeatherClientInterface,
) *GetWeatherUseCase {
	return &GetWeatherUseCase{
		CEPClient:     CEPClient,
		WeatherClient: WeatherClient,
	}
}

func (gw *GetWeatherUseCase) Execute(input GetWeatherRequest) (GetWeatherResponse, error) {
	address, err := gw.CEPClient.GetAddress(input.CEP)
	if err != nil {
		return GetWeatherResponse{}, errors.New("can not find zipcode")
	}

	climate, err := gw.WeatherClient.GetWeather(address.City)
	if err != nil {
		return GetWeatherResponse{}, errors.New("can not find weather for zipcode")
	}

	weather := entity.NewWeather(climate.Celsius)

	return GetWeatherResponse{
		Celsius:    weather.Celsius,
		Fahrenheit: weather.Fahrenheit,
		Kelvin:     weather.Kelvin,
	}, nil
}

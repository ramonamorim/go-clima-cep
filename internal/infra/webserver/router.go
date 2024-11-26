package webserver

import (
	"net/http"

	"github.com/ramonamorim/go-clima-cep/internal/infra/client"
	"github.com/ramonamorim/go-clima-cep/internal/usecase"
)

func Router(ApiKey string) http.Handler {
	cepClient := client.NewCepClient()
	weatherClient := client.NewWeatherClient(ApiKey)
	getWeatherUseCase := usecase.NewGetWeatherUseCase(cepClient, weatherClient)
	weather_controller := NewGetWeatherController(getWeatherUseCase)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /weather/{cep}", weather_controller.getWeather)

	return mux
}

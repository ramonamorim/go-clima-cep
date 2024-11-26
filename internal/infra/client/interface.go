package client

type CepResponse struct {
	Street       string `json:"street"`
	Complement   string `json:"complement"`
	Unit         string `json:"unit"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

type WeatherResponse struct {
	Celsius   float64 `json:"celsius"`
	Condition string  `json:"condition"`
}

type CepClientInterface interface {
	GetAddress(cep string) (*CepResponse, error)
}

type WeatherClientInterface interface {
	GetWeather(city string) (*WeatherResponse, error)
}

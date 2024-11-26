package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WeatherAPIKey string
	Port          string
}

var config *Config

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, tentando carregar do ambiente.")
	}

	config = &Config{
		WeatherAPIKey: os.Getenv("OPEN_WEATHERMAP_API_KEY"),
	}

	if config.WeatherAPIKey == "" {
		log.Fatal("Chave da API do WeatherAPI não configurada (OPEN_WEATHERMAP_API_KEY)")
	}
}

func GetConfig() *Config {
	if config == nil {
		log.Fatal("Configurações não carregadas. Certifique-se de chamar LoadConfig antes de GetConfig.")
	}
	return config
}

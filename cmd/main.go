package main

import (
	"log/slog"
	"net/http"

	"github.com/ramonamorim/go-clima-cep/config"
	"github.com/ramonamorim/go-clima-cep/internal/infra/webserver"
)

func main() {
	config.LoadConfig()
	cfg := config.GetConfig()

	r := webserver.Router(cfg.WeatherAPIKey)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	slog.Info("Server is running!")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

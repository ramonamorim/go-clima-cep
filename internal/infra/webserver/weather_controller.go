package webserver

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ramonamorim/go-clima-cep/internal/usecase"
	"github.com/ramonamorim/go-clima-cep/pkg/validator"
)

type GetWeatherController struct {
	GetWeatherUsecase usecase.GetWeatherUseCase
}

func NewGetWeatherController(GetWeatherUsecase *usecase.GetWeatherUseCase) *GetWeatherController {
	return &GetWeatherController{GetWeatherUsecase: *GetWeatherUsecase}
}

func (gw *GetWeatherController) getWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")

	if !validator.IsValidCEP(cep) {
		http.Error(w, errors.New("invalid zipcode").Error(), http.StatusUnprocessableEntity)
		return
	}

	req := usecase.GetWeatherRequest{CEP: cep}
	res, err := gw.GetWeatherUsecase.Execute(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

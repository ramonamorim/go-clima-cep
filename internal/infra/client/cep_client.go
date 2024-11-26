package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type CepClient struct{}

func NewCepClient() *CepClient {
	return &CepClient{}
}

type CepDTO struct {
	Erro        bool   `json:"erro"`
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (c CepClient) GetAddress(cep string) (*CepResponse, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data CepDTO
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Erro {
		return nil, errors.New("third party client error")
	}

	return &CepResponse{
		Street:       data.Logradouro,
		Complement:   data.Complemento,
		Unit:         data.Unidade,
		Neighborhood: data.Bairro,
		City:         data.Localidade,
		State:        data.UF,
	}, nil
}

package entity

import (
	"fmt"

	"github.com/ramonamorim/go-clima-cep/pkg/temperatura"
)

type Clima struct {
	Celsius    string
	Fahrenheit string
	Kelvin     string
}

func NewWeather(c float64) *Clima {
	f := temperatura.CelsiusToFahrenheit(c)
	k := temperatura.CelsiusToKelvin(c)

	celsius := fmt.Sprintf("%.2f", c)
	fahrenheit := fmt.Sprintf("%.2f", f)
	kelvin := fmt.Sprintf("%.2f", k)

	return &Clima{
		Celsius:    celsius,
		Fahrenheit: fahrenheit,
		Kelvin:     kelvin,
	}
}

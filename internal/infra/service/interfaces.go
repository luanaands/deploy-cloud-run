package service

import (
	"github.com/luanaands/deploy-cloud-run/internal/dto"
	"github.com/luanaands/deploy-cloud-run/internal/entity"
)

type CepInterface interface {
	GetViaCep(cep string, url string) (*dto.CepResponse, error)
}

type WeatherInterface interface {
	GetWeather(city string, apiKey string, baseURL string) (*entity.WeatherResponse, error)
}

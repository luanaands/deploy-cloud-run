package dto

import "github.com/luanaands/deploy-cloud-run/internal/entity"

func FromViaCep(resp *entity.CepViaCepResponse) *CepResponse {
	return &CepResponse{
		Localidade: resp.Localidade,
	}
}

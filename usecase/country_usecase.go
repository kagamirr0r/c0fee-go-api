package usecase

import (
	"c0fee-api/model"
	"c0fee-api/repository"
)

type ICountryUsecase interface {
	List() (model.CountriesResponse, error)
}

type countryUsecase struct {
	cr repository.ICountryRepository
}

func (cu *countryUsecase) List() (model.CountriesResponse, error) {
	countries := []model.Country{}
	err := cu.cr.List(&countries)
	if err != nil {
		return model.CountriesResponse{}, err
	}

	countryResponses := make([]model.CountryResponse, len(countries))
	for i, country := range countries {
		countryResponses[i] = model.CountryResponse{
			ID:    country.ID,
			Name:  country.Name,
			Code:  country.Code,
			Areas: country.Areas,
		}
	}

	return model.CountriesResponse{Countries: countryResponses, Count: uint(len(countries))}, nil
}

func NewCountryUsecase(cr repository.ICountryRepository) ICountryUsecase {
	return &countryUsecase{cr}
}

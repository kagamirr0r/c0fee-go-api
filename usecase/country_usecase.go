package usecase

import (
	"c0fee-api/model"
	"c0fee-api/repository"
)

type ICountryUsecase interface {
	Read(id uint) (model.CountryResponse, error)
	List() (model.CountriesResponse, error)
}

type countryUsecase struct {
	cr repository.ICountryRepository
}

func (cu *countryUsecase) Read(id uint) (model.CountryResponse, error) {
	storedCountry := model.Country{}
	if err := cu.cr.GetById(&storedCountry, id); err != nil {
		return model.CountryResponse{}, err
	}

	return storedCountry.ToResponse(), nil
}

func (cu *countryUsecase) List() (model.CountriesResponse, error) {
	countries := []model.Country{}
	err := cu.cr.List(&countries)
	if err != nil {
		return model.CountriesResponse{}, err
	}

	countryResponses := make([]model.CountryListResponse, len(countries))
	for i, country := range countries {
		countryResponses[i] = country.ToListResponse()
	}

	return model.CountriesResponse{Countries: countryResponses, Count: uint(len(countries))}, nil
}

func NewCountryUsecase(cr repository.ICountryRepository) ICountryUsecase {
	return &countryUsecase{cr}
}

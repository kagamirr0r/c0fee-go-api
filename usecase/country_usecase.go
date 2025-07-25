package usecase

import (
	"c0fee-api/dto"
	"c0fee-api/model"
	"c0fee-api/repository"
)

type ICountryUsecase interface {
	Read(id uint) (dto.CountryResponse, error)
	List() (dto.CountriesResponse, error)
}

type countryUsecase struct {
	cr repository.ICountryRepository
}

func (cu *countryUsecase) Read(id uint) (dto.CountryResponse, error) {
	storedCountry := model.Country{}
	if err := cu.cr.GetById(&storedCountry, id); err != nil {
		return dto.CountryResponse{}, err
	}

	return cu.convertToCountryResponse(&storedCountry), nil
}

func (cu *countryUsecase) List() (dto.CountriesResponse, error) {
	countries := []model.Country{}
	err := cu.cr.List(&countries)
	if err != nil {
		return dto.CountriesResponse{}, err
	}

	countryResponses := make([]dto.CountryListResponse, len(countries))
	for i, country := range countries {
		countryResponses[i] = dto.CountryListResponse{
			ID:   country.ID,
			Name: country.Name,
			Code: country.Code,
		}
	}

	return dto.CountriesResponse{Countries: countryResponses, Count: uint(len(countries))}, nil
}

func (cu *countryUsecase) convertToCountryResponse(country *model.Country) dto.CountryResponse {
	areas := make([]dto.AreaListResponse, len(country.Areas))
	for i, area := range country.Areas {
		areas[i] = dto.AreaListResponse{
			ID:   area.ID,
			Name: area.Name,
		}
	}

	return dto.CountryResponse{
		ID:    country.ID,
		Name:  country.Name,
		Code:  country.Code,
		Areas: areas,
	}
}

func NewCountryUsecase(cr repository.ICountryRepository) ICountryUsecase {
	return &countryUsecase{cr}
}

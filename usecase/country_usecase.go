package usecase

import (
	"c0fee-api/domain/country"
	"c0fee-api/dto"
)

type ICountryUsecase interface {
	Read(id uint) (dto.CountryOutput, error)
	List() (dto.CountriesOutput, error)
}

type countryUsecase struct {
	cr country.ICountryRepository
}

func (cu *countryUsecase) Read(id uint) (dto.CountryOutput, error) {
	var storedCountry country.Entity
	if err := cu.cr.GetById(&storedCountry, id); err != nil {
		return dto.CountryOutput{}, err
	}

	return cu.convertToCountryResponse(&storedCountry), nil
}

func (cu *countryUsecase) List() (dto.CountriesOutput, error) {
	var countries []country.Entity
	err := cu.cr.List(&countries)
	if err != nil {
		return dto.CountriesOutput{}, err
	}

	countryResponses := make([]dto.CountryListOutput, len(countries))
	for i, country := range countries {
		countryResponses[i] = dto.CountryListOutput{
			ID:   country.ID,
			Name: country.Name,
			Code: country.Code,
		}
	}

	return dto.CountriesOutput{Countries: countryResponses, Count: uint(len(countries))}, nil
}

func (cu *countryUsecase) convertToCountryResponse(countryEntity *country.Entity) dto.CountryOutput {
	areas := make([]dto.AreaListOutput, len(countryEntity.Areas))
	for i, area := range countryEntity.Areas {
		areas[i] = dto.AreaListOutput{
			ID:   area.ID,
			Name: area.Name,
		}
	}

	return dto.CountryOutput{
		ID:    countryEntity.ID,
		Name:  countryEntity.Name,
		Code:  countryEntity.Code,
		Areas: areas,
	}
}

func NewCountryUsecase(cr country.ICountryRepository) ICountryUsecase {
	return &countryUsecase{cr}
}

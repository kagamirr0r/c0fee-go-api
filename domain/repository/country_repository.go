package repository

import "c0fee-api/domain/entity"

type ICountryRepository interface {
	GetById(country *entity.Country, id uint) error
	List(countries *[]entity.Country) error
}
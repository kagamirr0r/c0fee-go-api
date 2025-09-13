package country

type ICountryRepository interface {
	GetById(country *Entity, id uint) error
	List(countries *[]Entity) error
}

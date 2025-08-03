package dto

// Output DTOs
type CountryOutput struct {
	ID    uint             `json:"id"`
	Name  string           `json:"name"`
	Code  string           `json:"code"`
	Areas []AreaListOutput `json:"areas"`
}

type CountryListOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type CountriesOutput struct {
	Countries []CountryListOutput `json:"countries"`
	Count     uint                `json:"count"`
}

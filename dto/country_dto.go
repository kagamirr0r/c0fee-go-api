package dto

// Response DTOs
type CountryResponse struct {
	ID    uint               `json:"id"`
	Name  string             `json:"name"`
	Code  string             `json:"code"`
	Areas []AreaListResponse `json:"areas"`
}

type CountryListResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type CountriesResponse struct {
	Countries []CountryListResponse `json:"countries"`
	Count     uint                  `json:"count"`
}

package dto

// Output DTOs
type AreaOutput struct {
	ID    uint             `json:"id"`
	Name  string           `json:"name"`
	Farms []FarmListOutput `json:"farms"`
}

type AreaListOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type AreasOutput struct {
	Areas []AreaListOutput `json:"areas"`
	Count uint             `json:"count"`
}

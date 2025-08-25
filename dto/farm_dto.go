package dto

// Output DTOs
type FarmOutput struct {
	ID      uint               `json:"id"`
	Name    string             `json:"name"`
	Farmers []FarmerListOutput `json:"farmers"`
}

type FarmListOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type FarmsOutput struct {
	Farms []FarmListOutput `json:"farms"`
	Count uint             `json:"count"`
}

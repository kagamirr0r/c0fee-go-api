package dto

// Response DTOs
type FarmResponse struct {
	ID      uint                 `json:"id"`
	Name    string               `json:"name"`
	Farmers []FarmerListResponse `json:"farmers"`
}

type FarmListResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type FarmsResponse struct {
	Farms []FarmListResponse `json:"farms"`
	Count uint               `json:"count"`
}

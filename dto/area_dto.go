package dto

// Response DTOs
type AreaResponse struct {
	ID    uint               `json:"id"`
	Name  string             `json:"name"`
	Farms []FarmListResponse `json:"farms"`
}

type AreaListResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type AreasResponse struct {
	Areas []AreaListResponse `json:"areas"`
	Count uint               `json:"count"`
}

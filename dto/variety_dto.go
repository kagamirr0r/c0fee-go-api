package dto

// Response DTOs
type VarietyListResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type VarietiesResponse struct {
	Varieties []VarietyListResponse `json:"varieties"`
	Count     uint                  `json:"count"`
}

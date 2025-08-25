package dto

// Output DTOs
type VarietyListOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type VarietiesOutput struct {
	Varieties []VarietyListOutput `json:"varieties"`
	Count     uint                `json:"count"`
}

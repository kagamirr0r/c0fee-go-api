package dto

// Output DTOs
type ProcessMethodOutput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ProcessMethodsOutput struct {
	ProcessMethods []ProcessMethodOutput `json:"process_methods"`
	Count          uint                  `json:"count"`
}

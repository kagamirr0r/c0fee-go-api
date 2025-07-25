package dto

// Response DTOs
type ProcessMethodResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ProcessMethodsResponse struct {
	ProcessMethods []ProcessMethodResponse `json:"process_methods"`
	Count          uint                    `json:"count"`
}

package dto

// Response DTOs
type RoasterResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	WebURL  string `json:"web_url"`
}

type RoastersResponse struct {
	Roasters []RoasterResponse `json:"roasters"`
	Count    uint              `json:"count"`
}

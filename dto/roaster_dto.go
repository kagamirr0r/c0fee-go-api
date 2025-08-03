package dto

// Output DTOs
type RoasterOutput struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	WebURL  string `json:"web_url"`
}

type RoastersOutput struct {
	Roasters []RoasterOutput `json:"roasters"`
	Count    uint            `json:"count"`
}

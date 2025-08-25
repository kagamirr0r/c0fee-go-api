package dto

// Output DTOs
type BeanRatingSummary struct {
	ID         uint          `json:"id"`
	User       IdNameSummary `json:"user"`
	Bitterness int           `json:"bitterness"`
	Acidity    int           `json:"acidity"`
	Body       int           `json:"body"`
	FlavorNote *string       `json:"flavor_note,omitempty"`
}

package dto

// Output DTOs
type BeanRatingOutput struct {
	ID         uint          `json:"id"`
	BeanID     uint          `json:"bean_id"`
	UserID     string        `json:"user_id"`
	User       IdNameSummary `json:"user"`
	Bitterness int           `json:"bitterness"`
	Acidity    int           `json:"acidity"`
	Body       int           `json:"body"`
	FlavorNote *string       `json:"flavor_note,omitempty"`
}

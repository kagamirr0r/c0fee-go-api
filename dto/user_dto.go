package dto

// Output DTOs
type UserOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

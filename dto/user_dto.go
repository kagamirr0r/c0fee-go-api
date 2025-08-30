package dto

import "github.com/google/uuid"

type UserInput struct {
	ID   uuid.UUID `json:"id" validate:"required,uuid4"`
	Name string    `json:"name" validate:"required"`
}

// Output DTOs
type UserOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

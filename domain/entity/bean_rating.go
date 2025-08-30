package entity

import (
	"time"

	"github.com/google/uuid"
)

type BeanRating struct {
	ID         uint
	BeanID     uint
	Bean       Bean
	UserID     uuid.UUID
	User       User
	Bitterness int
	Acidity    int
	Body       int
	FlavorNote string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
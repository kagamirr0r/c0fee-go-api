package bean_rating

import (
	"c0fee-api/domain/summary"
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID         uint
	BeanID     uint
	Bean       summary.Bean
	UserID     uuid.UUID
	User       summary.User
	Bitterness int
	Acidity    int
	Body       int
	FlavorNote string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Summary = summary.BeanRating

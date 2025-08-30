package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	AvatarKey string
	CreatedAt time.Time
	UpdatedAt time.Time
}
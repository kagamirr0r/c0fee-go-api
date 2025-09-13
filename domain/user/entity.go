package user

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        uuid.UUID
	Name      string
	AvatarKey string
	CreatedAt time.Time
	UpdatedAt time.Time
}

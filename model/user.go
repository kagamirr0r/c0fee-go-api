package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id" param:"id" gorm:"primary_key;type:uuid;" validate:"required"`
	Name      string         `json:"name" gorm:"unique" validate:"required"`
	AvatarKey string         `json:"avatar_key" gorm:"default:null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatar_url,omitempty"`
}

func (u *User) ToResponse(avatarURL string) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		AvatarURL: avatarURL,
	}
}

type UserBeansResponse struct {
	User  UserResponse   `json:"user"`
	Beans []BeanResponse `json:"beans"`
	Count uint           `json:"count"`
}

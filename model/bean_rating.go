package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BeanRating struct {
	ID         uint           `json:"id" param:"id" gorm:"primary_key;" validate:"required"`
	BeanID     uint           `json:"bean_id" gorm:"not null" validate:"required"`
	Bean       Bean           `json:"bean"`
	UserID     uuid.UUID      `json:"user_id" gorm:"type:uuid;not null" validate:"required"`
	User       User           `json:"user"`
	Bitterness int            `json:"bitterness" gorm:"comment:苦味の評価" validate:"required"`
	Acidity    int            `json:"acidity" gorm:"comment:酸味の評価" validate:"required"`
	Body       int            `json:"body" gorm:"comment:コク（ボディ）の評価" validate:"required"`
	FlavorNote string         `json:"flavor_note" gorm:"comment:フレーバーノート"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// type BeanRatingResponse struct {
// 	ID         uint      `json:"id"`
// 	BeanID     uint      `json:"bean_id"`
// 	UserID     uuid.UUID `json:"user_id"`
// 	User       User      `json:"user"`
// 	Bitterness int       `json:"bitterness"`
// 	Acidity    int       `json:"acidity"`
// 	Body       int       `json:"body"`
// }

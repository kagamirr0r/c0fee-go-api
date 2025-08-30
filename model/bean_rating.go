package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BeanRating struct {
	ID         uint `param:"id" gorm:"primary_key;"`
	BeanID     uint `gorm:"not null;uniqueIndex:idx_bean_user"`
	Bean       Bean
	UserID     uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_bean_user"`
	User       User
	Bitterness int    `gorm:"comment:苦味の評価"`
	Acidity    int    `gorm:"comment:酸味の評価"`
	Body       int    `gorm:"comment:コク（ボディ）の評価"`
	FlavorNote string `gorm:"comment:フレーバーノート"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
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

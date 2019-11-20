package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type (
	Tweet struct {
		ID     string `json:"id"`
		Text   string `json:"text"`
		UserID uint64 `json:"user_id"`
		CreatedAt  time.Time  `json:"created_at"`
		UpdatedAt  time.Time  `json:"updated_at"`
		DeletedAt  *time.Time `json:"deleted_at"`
	}
)

func (t *Tweet) Create(db *gorm.DB, input CreateTweetInput, user *User) error {
	t.Text = input.Text
	t.UserID = user.ID
	return db.Create(&t).Error
}

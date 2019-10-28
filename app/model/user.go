package model

import "time"
import "github.com/jinzhu/gorm"

type (
	User struct {
		Name      string     `json:"name"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}
)

func (u *User) Create(db *gorm.DB, input CreateUserInput) error {
	u.Name = input.Name
	return db.Create(u).Error
}

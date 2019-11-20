package model

import (
	"errors"
	"strconv"
	"time"
)
import "github.com/jinzhu/gorm"

type (
	User struct {
		ID         uint64     `json:"id"`
		Identifier string     `json:"identifier"`
		Name       string     `json:"name"`
		CreatedAt  time.Time  `json:"created_at"`
		UpdatedAt  time.Time  `json:"updated_at"`
		DeletedAt  *time.Time `json:"deleted_at"`
	}
)

func (u *User) Create(db *gorm.DB, input CreateUserInput) error {
	u.Name = input.Name
	u.Identifier = input.Identifier
	return db.Create(u).Error
}

func (u *User) FindByIdentifier(db *gorm.DB, id string) error {
	if id == "" {
		return errors.New("id_not_found")
	}

	uID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return err
	}
	return db.Where("id = ?", uID).First(&u).Error
}

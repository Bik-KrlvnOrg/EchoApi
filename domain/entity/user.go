package entity

import (
	"encoding/json"
	"io"
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;auto_increment" json:"id" `
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []*User

func (users *Users) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(users)
}

func (user *User) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(user)
}

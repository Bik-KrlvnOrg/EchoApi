package dto

import (
	"encoding/json"
	"io"
)

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

func (user *User) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(user)
}

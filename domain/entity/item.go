package entity

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"io"
)

type JSONB map[string]interface{}

type Item struct {
	gorm.Model
	Echos JSONB `sql:"type:jsonb"`
}

type Items []*Item

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

func (item *Item) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(item)
}

func (items *Items) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(items)
}

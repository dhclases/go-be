package phonebook

import (
	"github.com/jinzhu/gorm"
)

// PhoneBooks - maping table phone_books
type PhoneBooks struct {
	gorm.Model
	Name        string
	PhoneNumber string
}

// Phone - struct for json
type Phone struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
}

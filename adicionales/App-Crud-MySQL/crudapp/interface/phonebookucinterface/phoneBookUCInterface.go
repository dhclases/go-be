package phonebookucinterface

import (
	phoneModel "crudgolang/model/phonebook"
)

// PhoneBookUCI - phone book usecase interface
type PhoneBookUCI interface {
	InsertData(string, string) error
	EditData(int64, string, string) error
	DeleteData(int64) error
	GetAll() ([]phoneModel.Phone, error)
	GetByID(int64) (phoneModel.Phone, error)
}

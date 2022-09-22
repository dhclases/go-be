package service

import (
	phoneI "crudgolang/interface/phonebookinterface"
	phoneUCI "crudgolang/interface/phonebookucinterface"

	phoneModel "crudgolang/model/phonebook"
)

// phoneBookService - implement PhoneBookInterface
type phoneBookService struct {
	phoneI.PhoneBookI
}

// NewphoneBookService - initial phone book use case
func NewPhoneBookService(phoneInterface phoneI.PhoneBookI) phoneUCI.PhoneBookUCI {
	return &phoneBookService{
		phoneInterface,
	}
}

// GetAll - get data list from table phone_books
func (pbs *phoneBookService) GetAll() ([]phoneModel.Phone, error) {
	var (
		list    []phoneModel.PhoneBooks
		myDatas []phoneModel.Phone
	)

	list, err := pbs.PhoneBookI.FindAll()
	if err != nil {
		return myDatas, err
	}

	for _, val := range list {
		var detail phoneModel.Phone
		detail.ID = val.ID
		detail.Name = val.Name
		detail.PhoneNumber = val.PhoneNumber
		myDatas = append(myDatas, detail)
	}

	return myDatas, nil
}

// GetByID - get data from table phone_books by id
func (pbs *phoneBookService) GetByID(id int64) (phoneModel.Phone, error) {
	var (
		data   phoneModel.PhoneBooks
		myData phoneModel.Phone
	)

	data, err := pbs.PhoneBookI.FindByID(id)
	if err != nil {
		return myData, err
	}

	myData.ID = data.ID
	myData.Name = data.Name
	myData.PhoneNumber = data.PhoneNumber

	return myData, nil
}

// InsertData - insert data into table phone_books
func (pbs *phoneBookService) InsertData(name, phone string) error {
	err := pbs.PhoneBookI.CreateData(name, phone)

	if err != nil {
		return err
	}

	return nil
}

// EditData - edit data from table phone_books by id
func (pbs *phoneBookService) EditData(id int64, name, phone string) error {
	err := pbs.PhoneBookI.UpdateData(id, name, phone)

	if err != nil {
		return err
	}

	return nil
}

// DeleteData - delete data from table phone_books by id
func (pbs *phoneBookService) DeleteData(id int64) error {
	err := pbs.PhoneBookI.DeleteData(id)
	if err != nil {
		return err
	}

	return nil
}

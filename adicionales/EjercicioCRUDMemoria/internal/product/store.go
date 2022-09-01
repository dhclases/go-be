package product

import (
	"actividad/clase/internal/domain"
	"encoding/json"
	"os"
)

func LoadProducts(path string) []domain.Product {
	list := []domain.Product{}

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}

	return list

}

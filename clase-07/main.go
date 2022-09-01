package main

import (
	"fmt"

	"github.com/pkg/errors"
)

const lowSalary = 10000

var ErrSalary = errors.New("salario muy bajo")

func validateSalary(salary int) error {
	if salary < lowSalary {
		// se usa el verbo %w para hacer wraps del error go 1.13
		return fmt.Errorf("validate input: %w", ErrSalary)
	}
	return nil
}

func main() {
	err := validateSalary(9000)

	if errors.Is(err, ErrSalary) {
		fmt.Println("salario bajo")
	} else {
		fmt.Println("salario dentro de lo esperado")
	}

}

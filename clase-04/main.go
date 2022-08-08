package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOFBirth string
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

// Definiendo el receptor con value reciever
func (e Employee) PrintEmployee() {
	e.Position = "dfsdfsdf"
	fmt.Printf("Employee: %v \n", e)
}

func main() {
	person1 := Person{
		ID:          1,
		Name:        "John",
		DateOFBirth: "2013-01-01T00:00:00Z",
	}

	employee1 := Employee{
		ID:       2,
		Position: "Developer",
		Person:   person1,
	}

	employee1.PrintEmployee()

	fmt.Printf("Employee out scope: %v \n", employee1)
}

/*

 Value reciever vs Pointer receiver

 El receptor de valor hace una copia del tipo y lo pasa a la función.
 La pila de funciones ahora contiene un objeto igual
 pero en una ubicación diferente en la memoria.
 Eso significa que cualquier cambio realizado en el objeto pasado permanecerá local
 para el método. El objeto original permanecerá sin cambios.

El receptor de puntero pasa la dirección de un tipo a la función.
La pila de funciones tiene una referencia al objeto original.
Entonces, cualquier modificación en el objeto pasado modificará el objeto original.


*/

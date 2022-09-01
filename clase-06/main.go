package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Gender string
	Age    int
	Family []Person `json:",omitempty"`
}

type humano struct {
	Nombre   string
	Apellido string
}

func main() {
	p1 := Person{
		Name:   "John",
		Gender: "Male",
		Age:    18,
		Family: []Person{
			{
				Name:   "John II",
				Gender: "Male",
				Age:    8,
			},
			{
				Name:   "Rose",
				Gender: "Female",
				Age:    2,
			},
		},
	}

	fmt.Printf("valor de p1: %v \n", p1)
	fmt.Printf("tipo de p1: %T \n", p1)

	json, _ := json.Marshal(p1)
	fmt.Println(string(json))

	fmt.Println()
	fmt.Println("***")
	fmt.Println()

	personita := &humano{"Pepe", "Morales"}

	fmt.Printf("valor de personita: %v \n", personita)
	fmt.Printf("Tipo personita: %T \n", personita)
	fmt.Println("Struct valor *personita :", *personita)
	fmt.Println("Direccion memoria &(*personita) :", &(*personita))
	fmt.Println("Valor del struct *(&(*personita)) :", *(&(*personita)))
	fmt.Println("Valor del campo Nombre  (*(&(*personita)) :", (*(&(*personita))).Nombre)

	fmt.Println()
	fmt.Println("***")
	fmt.Println()

	var pointer *int // Forma 1

	// Forma 2
	var v int
	pointer = &v

	// Forma 3
	pointer2 := new(int)

	/*
	   -- Wrong
	   var v2 int
	   pointer := *v2
	*/

	fmt.Printf("tipo de pointer: %T \n", pointer)
	fmt.Printf("tipo de v: %T \n", v)
	fmt.Printf("tipo de pointer2: %T \n", pointer2)

	fmt.Println()
	fmt.Println("***")
	fmt.Println()

	var newVariable []interface{}

	newVariable = append(newVariable, 12, "test", 1.20, -1.36, personita)

	for i, v := range newVariable {
		fmt.Printf("Key = %v\t valor = %v\t tipo= %T \n", i, v, v)
	}

}

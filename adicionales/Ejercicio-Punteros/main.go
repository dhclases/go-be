package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		printHeader("*** Definir puntero ***")
		definirPuntero()
	*/
	/*
		printHeader("*** Inicializar puntero ***")
		inicializarPuntero()

		printHeader("*** Desreferenciar el puntero ***")
		DesreferenciarPuntero()


						printHeader("*** Cambiar el valor atraves del puntero ***")
						cambiarValorPorPuntero()


					printHeader("*** Puntero de Puntero ***")
					punteroDePuntero()
				printHeader("*** Comparar Punteros ***")
				compararPunteros()


			printHeader("*** Struct por referencia ***")
	*/
	p := Person{
		Name:        "John",
		Age:         21,
		DateOfBirth: time.Date(1980, 11, 17, 20, 34, 58, 651387237, time.UTC),
	}

	printHeader("*** Struct por valor ***")
	fmt.Println("person antes de llamar funcion inmutable: ", p)
	p.inmutablePerson()
	fmt.Println("person despues de llamar funcion inmutable: ", p)

	printHeader("*** Struct por Referencia ***")
	fmt.Println("person antes de llamar funcion mutable: ", p)
	p.mutablePerson()
	fmt.Println("person despues de llamar funcion mutable: ", p)

}

func definirPuntero() {
	// Esquema var name *T
	// Punto clave usar *
	var p *int
	fmt.Println("p = ", p)
}

func inicializarPuntero() {
	// Punto clave usar * y &
	var a = 5.67
	var p2 = &a

	fmt.Println("Valor almacenado en a = ", a)
	fmt.Println("Address de  a = ", &a)
	fmt.Println("Valor en p2 = ", p2)
}

func DesreferenciarPuntero() {
	var a3 = 100
	var p3 = &a3

	fmt.Println("a3 = ", a3)
	fmt.Println("p3 = ", p3)
	fmt.Println("*p3 = ", *p3)
}

func cambiarValorPorPuntero() {
	var a4 = 1000
	var p4 = &a4

	fmt.Println("a4 (antes) = ", a4)

	// Cambiando el valor de la variable por atraves del puntero
	*p4 = 2000

	fmt.Println("a4 (despues) = ", a4)

}

func punteroDePuntero() {
	var a = 7.98
	var p = &a
	var pp = &p

	fmt.Println("a = ", a)
	fmt.Println("address of a = ", &a)

	fmt.Println("p = ", p)
	fmt.Println("address of p = ", &p)

	fmt.Println("pp = ", pp)

	// Dereferencing a pointer to pointer
	fmt.Println("*pp = ", *pp)
	fmt.Println("**pp = ", **pp)
}

func compararPunteros() {
	var a = 75
	var p1 = &a
	var p2 = &a

	if p1 == p2 {
		fmt.Println("p1 y p2 apunta a la misma direccion")
	}
}

func printHeader(s string) {
	fmt.Println()
	fmt.Println(s)
	fmt.Println()

}

type Person struct {
	Name        string
	Age         int
	DateOfBirth time.Time
}

func (p *Person) mutablePerson() {

	p.Name = "Messi"
	p.Age = 36
	fmt.Printf("Name:%s Age:%d \n", p.Name, p.Age)
	fmt.Println()
}

func (p Person) inmutablePerson() {
	p.Age = 3
	p.Name = "default3"
	fmt.Printf("Name:%s Age:%d \n", p.Name, p.Age)
	fmt.Println()
}

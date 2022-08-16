package main

import "fmt"

/*
Objetivo ver como usar defer
*/
func main() {
	// Uso de defer
	deferExample1()

	// Varios defer y su invocaciones
	deferExample2()

	// uso de defer para recuperarno de un panic
	deferExample3()

	// uso de defer con funciones anónimas con o sin parametros
	deferExample4()

}

/*
En este ejemplo de muestra como usar el defer o en otras palabras al cerrar una function
ejecutar un codigo
*/
func deferExample1() {
	defer fmt.Println("Esto se imprime luego de cerrar deferExample1")
	fmt.Println("Ejecutando funcion defer Example1")
}

/*
En este ejemplo de muestra como usar varia funciones defer
Varias defer se juntas se ejecutan como en cola, la primera entrar la ultima en salir
En otras palabras, varias invocaciones defer se ejecutan de la ultima primero, luego la penultima
y asi sucesivbamente
*/
func deferExample2() {
	defer fmt.Println("Primer defer")
	defer fmt.Println("Segundo defer")
	defer fmt.Println("Tercer defer")

	fmt.Println()
	fmt.Println("*** Ejecutando defer Example 2")
}

/*
   Defer tienes muchos usos, uno de ellos es recuperarnos de panic
   En este ejemplo si firstName y lastName son nil disparamos un panic y nos recuperamos
   suavemente del mismo
*/

func revoverWithMessage() {

	/*
	   Go nos provee de la función recover para recuperarnos de un panic o gorutine
	*/
	if r := recover(); r != nil {
		fmt.Println("Nos recuperamos desde", r)
	}
}

func fullName(firstname, lastName string) string {
	// Si ocurre un panic al cerrar la funcion nos recuperamos suavemente
	defer revoverWithMessage()
	if firstname == "" {
		panic("firstname esta vacçioo")
	}
	if lastName == "" {
		panic("lastname esta vacío")
	}

	return fmt.Sprintf("%s %s \n", firstname, lastName)
}

func deferExample3() {
	firstName := "joe"
	lastName := "Terry"

	fmt.Println()
	fmt.Println("*** Ejecutando defer Example 3")

	// Llamamos la funcion con los valores correctos
	fmt.Println(fullName(firstName, lastName))

	// Llamamos la funcion con los valores INCORRECTOS
	fmt.Println(fullName("", ""))

}

func deferExample4() {
	// 1er defer sin parametros y por funcion anónima
	defer func() {
		fmt.Println(" Defer dentro de la funcion anónima sin párametro")
	}()

	// 2do defer sin parametros, por funcion anónima y con parametros
	defer func(valor string) {
		fmt.Println(" Defer dentro de la funcion anónima con párametro; ", valor)
	}("Mi valor")

	// Ejecucion normal
	fmt.Println()
	fmt.Println("*** Ejecutando defer Example 4")

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
   Tema 0 : Ejeuctar varias tareas
*/

// Funcion utilitaria para dormir el proceso en un rango de 100 a 300 ms
func goToSleep() {
	rand.Seed(time.Now().UnixNano())
	min := 100
	max := 300
	t := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(t) * time.Millisecond)
}

// Tarea 1: Mostrar numeros
func showNumber() {
	count := 20
	for i := 0; i < count; i++ {
		fmt.Printf("Number: %d\n", i)
	}
	goToSleep()
}

// Tarea 2: Mostrar Letras
func showAlphabets() {
	for j := 'a'; j <= 'z'; j++ {
		fmt.Println("value of j=", string(j))
	}
	goToSleep()
}

func main() {

	// Tomar el tiemppo inicial
	start := time.Now()

	// Ejecutar las tareas 1 y 2 30 veces
	for i := 0; i < 30; i++ {
		showNumber()
		showAlphabets()

	}

	// Medir tiempo de ejecucion
	elapsed := time.Since(start)

	fmt.Printf("Process took: %v\n", elapsed)

}

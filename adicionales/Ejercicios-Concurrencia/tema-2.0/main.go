package main

import (
	"fmt"
)

var total int = 0

// Ejemplo condicion de carrera
// Use gor run -race main.go si hay una condicion de carrera

func main() {

	for i := 0; i < 30; i++ {
		go increment()
		go increment()
	}

	fmt.Printf("total: %d\n", total)
}

func increment() {

	total++

}

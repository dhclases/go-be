package main

import (
	"fmt"
	"time"
)

func main() {

	// Creamos un canal del tipo entero
	ch := make(chan int)

	// iniciamos una funcion anónima
	go func() {
		time.Sleep(10 * time.Millisecond)
		// Enviamos un numero (42)
		ch <- 42
		// Dormimos el canal unos minutos
		time.Sleep(10 * time.Millisecond)

	}()

	// Leemos lo que estan en el canal
	fmt.Println(<-ch)

}

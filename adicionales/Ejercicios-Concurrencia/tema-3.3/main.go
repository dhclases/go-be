package main

import (
	"fmt"
	"sync"
)

// Ejemplo de tipos de canales

func main() {
	wg := &sync.WaitGroup{}
	// Canal de Buffer (2 items)
	ch := make(chan int, 1)

	wg.Add(2)
	// Go rutina de solo recepcion
	//         <_chan int indica solo recibo
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println("Activando Recepcion")
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	// Go rutina de solo recepcion
	//         chan<- int indica solo envioo
	go func(ch chan<- int, wg *sync.WaitGroup) {
		fmt.Println("Activando Emision")
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch, wg)

	wg.Wait()

	fmt.Printf("Ultimo item:%v", <-ch)
}

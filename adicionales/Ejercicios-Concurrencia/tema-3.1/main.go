package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)

	// Go rutina bidireccional de recepcion
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("Activando Recepcion")
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	// Go rutina direcional de emision
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("Activando Emision")
		ch <- 42
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

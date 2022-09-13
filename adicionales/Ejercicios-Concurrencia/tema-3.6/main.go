package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	// Recepcion de lotes de enteros
	go func(ch <-chan int, wg *sync.WaitGroup) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch, wg)
	// Emision de lotes de enteros
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

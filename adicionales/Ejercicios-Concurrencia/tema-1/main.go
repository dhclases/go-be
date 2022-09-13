package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Funcion utilitaria para dormir el proceso en un rango de 100 a 300 ms
func goToSleep() {
	rand.Seed(time.Now().UnixNano())
	min := 100
	max := 300
	t := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(t) * time.Millisecond)
}

/*
   Tema 1 : Go Rutina
*/

func showNumber(wg *sync.WaitGroup) {
	count := 20
	for i := 0; i < count; i++ {
		fmt.Printf("Number: %d\n", i)
	}
	goToSleep()
	wg.Done()
}

func showAlphabets(wg *sync.WaitGroup) {
	for j := 'a'; j <= 'z'; j++ {
		fmt.Println("value of j=", string(j))
	}
	goToSleep()
	wg.Done()
}

func main() {
	// Tomar el tiemppo inicial
	start := time.Now()
	wg := &sync.WaitGroup{}

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go showNumber(wg)
		wg.Add(1)
		go showAlphabets(wg)
	}

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("Process took: %v\n", elapsed)

}

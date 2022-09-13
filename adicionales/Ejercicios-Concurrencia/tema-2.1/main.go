package main

import (
	"fmt"
	"sync"
)

var total int = 0

// Ejemplo condicion de carrera
// Use gor run -race main.go si hay una condicion de carrera

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(60)
	for i := 0; i < 30; i++ {
		go increment(&wg, &m)
		go increment(&wg, &m)
	}
	wg.Wait()
	fmt.Printf("total: %d\n", total)
}

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	temp := total
	temp++
	total = temp
	m.Unlock()
	wg.Done()
}

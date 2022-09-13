package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"hilos.dh.com/model"
)

var cache = map[int]model.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan model.Book)
	dbCh := make(chan model.Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- model.Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- model.Book) {
			if b, ok := queryDatabase(id); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)

		go func(cacheCh, dbCh <-chan model.Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("*** from cache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("*** from database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (model.Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int) (model.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range model.Books {
		if b.ID == id {
			return b, true
		}
	}

	return model.Book{}, false
}

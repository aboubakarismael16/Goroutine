package main

import (
	"fmt"
	"sync"
)

var x int

func increment(wg *sync.WaitGroup, m *sync.Mutex)  {
	m.Lock()
	x += 1
	m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg,&m)
	}

	wg.Wait()
	fmt.Println("the final value of x1 : ", x)
}

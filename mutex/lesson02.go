package main

import (
	"fmt"
	"sync"
)

var x1 = 0

func increment1(wg *sync.WaitGroup, ch chan bool)  {
	ch <- true
	x1 = x1 + 1
	<- ch
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment1(&wg, ch1)
	}

	wg.Wait()
	fmt.Println("final value of x1 : ", x1)

}

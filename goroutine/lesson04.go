package main

import (
	"fmt"
	"sync"
)


func main() {
	var wg sync.WaitGroup
	var x, y int
	wg.Add(2)
	x, y = 1, 2

	go func() {
		defer wg.Done()
		fmt.Println(x)
	}()
	go func() {
		defer wg.Done()
		fmt.Println(y)
	}()

	wg.Wait()
	fmt.Println("main function")
	//time.Sleep(1*time.Second)

}

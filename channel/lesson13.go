package main

import (
	"fmt"
	"sync"
)

var b chan int
var wg13 sync.WaitGroup

func noBufChannel()  {
	fmt.Println(b) // nil
	b = make(chan int) // you need to initialize the channel
	wg13.Add(1)
	go func() {
		defer wg13.Done()
		x := <- b
		fmt.Println("Received the sent value from b : ",x)
	}()

	b <- 10
	fmt.Println("noBufChannel : 10 send to channel b ...")
	wg13.Wait()
}

func bufChannel()  {
	fmt.Println(b) // nil
	b = make(chan int, 10) // you need to initialize the channel
	b <- 10
	fmt.Println("bufChannel : 10 send to channel b ...")
	b <- 20
	fmt.Println("bufChannel : 20 send to channel b ...")
	x := <- b
	fmt.Println("bufChannel :Received the sent value from b ",x)
	close(b)

}

func main() {
	bufChannel()
}

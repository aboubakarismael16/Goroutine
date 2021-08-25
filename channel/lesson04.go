package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest(r chan<- int32)   {

	time.Sleep(3 * time.Millisecond)
	r <- rand.Int31n(100)

}

func square(a,b int32) int32  {
	return a*a + b*b
}

func main() {

	rand.Seed(time.Now().UnixNano())

	//a, b := make(chan int32), make(chan int32)
	//go longTimeRequest(a)
	//go longTimeRequest(b)


	results := make(chan int32, 2)  // The channel can be buffered or not.
	go longTimeRequest(results)
	go longTimeRequest(results)

	fmt.Println(square(<-results,<-results))
}

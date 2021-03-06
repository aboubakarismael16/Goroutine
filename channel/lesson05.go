package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest2() <-chan int32 {
	r := make(chan int32)

	go func() {
		// Simulate a workload.
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a, b := longTimeRequest2(), longTimeRequest2()
	fmt.Println(sumSquares(<-a, <-b))
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source(c chan <- int32)  {
	ra, rb := rand.Int31(), rand.Intn(3)*1
	time.Sleep(time.Duration(rb)* time.Second) // sleep 1s/2s/3s
	c <- ra
}

func main() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()

	c := make(chan int32, 5) // must be buffered channel
	for i := 0; i < cap(c); i++ {
		go source(c)
	}

	rnd := <- c // only the first response will be used
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)

}

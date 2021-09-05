package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(5*time.Second)
		ch1 <- 200
	}()

	go func() {
		time.Sleep(6*time.Second)
		ch2 <- 100
	}()

	select {
	case num1 := <- ch1:
		fmt.Println("already receive ch1 value :", num1)
	case num2, ok := <- ch2:
		if ok {
			fmt.Println("already receive ch2 value :", num2)
		} else {
			fmt.Println("ch2 closed")
		}
	case <- time.After(4*time.Second):
		fmt.Println("timeout")


	}
}

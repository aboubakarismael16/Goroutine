package main

import "fmt"

func main()  {

	ch1 := make(chan int)     // unbuffered
	fmt.Println(len(ch1), cap(ch1))  //0 0
	ch1 <- 10  // unbuffered channel, it's need to another goroutine to lock the block

	ch2 := make(chan int, 5) // buffered , cap = 5
	fmt.Println(len(ch2), cap(ch2))

	for i := 0; i < 50; i++ {
		ch2 <- i+10
	}
	//ch2 <- 10
	//ch2 <- 20
	//ch2 <- 30
	//ch2 <- 40
	//ch2 <- 50
	fmt.Println("------------")
	fmt.Println(len(ch2))

}
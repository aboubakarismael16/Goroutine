package main

import "fmt"

func main()  {
	ch1 := make(chan int, 5)
	ch1 <- 10
	fmt.Println(len(ch1), cap(ch1))
	ch1 <- 20
	ch1 <- 30
	ch1 <- 40
	ch1 <- 50
	fmt.Println("------------")
	fmt.Println(len(ch1))

}
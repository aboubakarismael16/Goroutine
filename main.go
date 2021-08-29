package main

import (
	"fmt"
)

func main()  {

	ch := make(chan string)
	done := make(chan bool)
	go sendData3(ch, done)
	data := <-ch
	fmt.Println("child goroutine zhuan lai : ", data)
	ch <- "I am main"

	<-done

	fmt.Println("main function")

}

func sendData3(ch chan string, done chan bool)  {
	ch <- "I am xiao ming"
	data := <- ch
	fmt.Println("main goroutine zhuan lai: ", data)
	done <- true
}
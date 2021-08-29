package main

import (
	"fmt"
	"strconv"
	"time"
)

func main()  {

	ch2 := make(chan string, 4)
	go sendData3(ch2)
	for  {
		time.Sleep(1*time.Second)
		v,ok := <- ch2
		if !ok {
			fmt.Println("Finished to read", ok)
			break
		}

		fmt.Println("\tthe read value is :", v)
	}

	fmt.Println("main function")

}

func sendData3(ch chan string)  {
	for i := 0; i < 10; i++ {
		ch <- "Data" + " " + strconv.Itoa(i)
		fmt.Println("child goroutine ,", i)
	}
	close(ch)
}
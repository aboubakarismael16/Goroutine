package main

import "fmt"

func main()  {
	var ch1 chan bool
	fmt.Println(ch1)
	ch1 = make(chan bool)
	fmt.Println(ch1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("start to print value of i:", i)
		}
		ch1 <- true
		fmt.Println("end to send")
	}()

	data := <- ch1
	fmt.Println("data-->", data)
	fmt.Println("main goroutine over")
}
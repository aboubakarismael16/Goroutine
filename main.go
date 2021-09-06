package main

import (
	"fmt"
	"time"
)

func main()  {
	go func() {
		fmt.Println("goroutine_1")
	}()

	go func() {
		fmt.Println("goroutine_2")
	}()

	func(){
		fmt.Println("func_1")
	}()
	time.Sleep(time.Second)
}
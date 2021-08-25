package main

import (
	"fmt"
	"time"
)

var s = "hello world"

func numbers()  {
	for i := 0; i < 10; i++ {
		time.Sleep(250 *time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabet()  {
	for i := 'a'; i <= 'p' ; i++ {
		time.Sleep(500*time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main()  {
	go numbers()
	go alphabet()
	time.Sleep(3*time.Second)
	fmt.Println("main terminated")

}
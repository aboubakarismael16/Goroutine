package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var ticket = 10


func main() {
	go saleTicket("gate 1")
	go saleTicket("gate 2")
	go saleTicket("gate 3")
	go saleTicket("gate 4")
	time.Sleep(5*time.Second)
	fmt.Println(runtime.NumGoroutine())
}

func saleTicket(name string)  {
	rand.Seed(time.Now().UnixNano())
	for {
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "sold", ticket)
			ticket--
		} else {
			fmt.Println(name, "not ticket to sold")
			break
		}
	}
}
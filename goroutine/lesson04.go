package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var ticket = 10
var wg04 sync.WaitGroup
var mutex sync.Mutex

func main() {
	wg04.Add(4)
	go saleTicket("gate 1")
	go saleTicket("gate 2")
	go saleTicket("gate 3")
	go saleTicket("gate 4")
	//time.Sleep(5*time.Second)
	wg04.Wait()
	fmt.Println(runtime.NumGoroutine())
}

func saleTicket(name string)  {
	rand.Seed(time.Now().UnixNano())
	defer wg04.Done()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "sold", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println(name, "not ticket to sold")
			break
		}
		mutex.Unlock()
	}
}
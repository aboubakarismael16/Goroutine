package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}


func player(name string, table chan *Ball)  {
	for {
		ball := <- table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(time.Second)
		table <- ball
	}
}

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball)
	time.Sleep(time.Second)
	<- table
}

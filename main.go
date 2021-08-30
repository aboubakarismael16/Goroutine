package main

import (
	"fmt"
	"time"
)

func main()  {

	//timer := time.NewTimer(1*time.Second)
	//fmt.Printf("%T\n", timer)
	//fmt.Println(time.Now())
	//
	//ch := timer.C
	//fmt.Println(<-ch)

	fmt.Println("--------------")

	timer2 := time.NewTimer(5 * time.Second)

	go func() {
		 <- timer2.C

		 fmt.Println("timer2 finished")
	}()

	time.Sleep(3*time.Second)
	stop := timer2.Stop()

	if stop {
		fmt.Println("timer 2 stopped")
	}

}
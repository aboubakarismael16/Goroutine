package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int)  {
	for j := range jobs {
		fmt.Printf("Worker: %d start job: %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker: %d end: %d\n", id, j)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	//Run 3 goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	// giving 5 works
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	//receive the given work
	for r := 1; r <=5 ; r++ {
		<-  result
	}
}



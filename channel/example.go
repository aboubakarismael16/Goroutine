package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

var wgEx sync.WaitGroup

func takeZero(tz chan <- *job)  {
	defer wgEx.Done()
	for  {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		tz <- newJob
		time.Sleep(time.Millisecond*500)
	}
}

func printResult(tz <- chan *job, resultChan chan <- *result)  {
	defer wgEx.Done()
	for {
		job := <- tz
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n/10
		}

		newResult := &result{
			job: job,
			sum: sum,
		}

		resultChan <- newResult
	}
}

func main() {
	wgEx.Add(1)
	go takeZero(jobChan)
	wgEx.Add(24)
	for i := 0; i < 24; i++ {
		go printResult(jobChan, resultChan)
	}

	for result := range resultChan {
		fmt.Printf("value: %d sum: %d\n", result.job.value, result.sum)
	}

	wgEx.Wait()

}

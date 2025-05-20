package main

import (
	"fmt"
	"sync"
	"time"
)

type SharedCounter struct {
	mu    sync.Mutex
	count int
}

func (s *SharedCounter) Increment() {
	s.mu.Lock()
	s.count++
	s.mu.Unlock()
}

func (s *SharedCounter) Value() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}

func worker(_ int, wg *sync.WaitGroup, resultChannel chan<- int, jobChannel <-chan int, counter *SharedCounter) {

	defer wg.Done()
	for jobs := range jobChannel {
		time.Sleep(time.Millisecond * 1000)

		resultChannel <- jobs * 2

		counter.Increment()
	}

}

func main() {
	const numWorkers = 3
	const numJobs = 10

	jobChannel := make(chan int, numJobs)
	resultChannel := make(chan int, numJobs)

	counter := SharedCounter{}

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, resultChannel, jobChannel, &counter)
	}

	for i := 0; i < numJobs; i++ {
		jobChannel <- i
	}

	close(jobChannel)

	wg.Wait()
	close(resultChannel)

	for val := range resultChannel {
		fmt.Println("Result: ", val)
	}

	fmt.Println("Shared counter value: ", counter.Value())

}

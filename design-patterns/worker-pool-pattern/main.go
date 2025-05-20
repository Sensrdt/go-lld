package main

import (
	"fmt"
	"time"
)

type WorkerFunc[T any, R any] func(T) R

func StartWorkerPool[T any, R any](numWorkers int, jobs <-chan T, worker WorkerFunc[T, R]) <-chan R {
	results := make(chan R)

	for i := 1; i <= numWorkers; i++ {
		go func(id int) {
			for job := range jobs {
				fmt.Printf("Worker %d processing job: %+v\n", id, job)
				result := worker(job)
				results <- result
			}
		}(i)
	}

	return results
}

func heavyJob(n int) int {
	time.Sleep(500 * time.Millisecond)
	return n * 2
}

func main() {
	const (
		numJobs    = 10
		numWorkers = 3
	)

	jobs := make(chan int, numJobs)

	go func() {
		defer close(jobs)
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
	}()

	results := StartWorkerPool(numWorkers, jobs, heavyJob)

	for i := 1; i <= numJobs; i++ {
		fmt.Println("Result:", <-results)
	}
}

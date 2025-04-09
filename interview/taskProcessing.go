package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task represents a unit of work
type Task struct {
	ID   int
	Data int // Some input data
}

// WorkerPool manages task execution
type WorkerPool struct {
	WorkerCount int
	TaskQueue   chan Task
	ResultQueue chan int
	wg          sync.WaitGroup
}

// ProcessTask simulates processing a task
func ProcessTask(task Task) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // Simulate work
	return task.Data * 2
}

// Worker listens for tasks and processes them
func (wp *WorkerPool) Worker(workerID int) {
	defer wp.wg.Done()

	for task := range wp.TaskQueue {
		fmt.Printf("Worker %d processing task %d\n", workerID, task.ID)
		result := ProcessTask(task)
		wp.ResultQueue <- result
	}
}

func (wp *WorkerPool) Start() {
	for i := 1; i <= wp.WorkerCount; i++ {
		wp.wg.Add(1)
		go wp.Worker(i)
	}
}

func (wp *WorkerPool) Shutdown() {
	wp.wg.Wait()
	close(wp.ResultQueue)
}

func WorkerPoolPattern() {
	const numWorkers = 3
	const numTasks = 10

	workerPool := WorkerPool{
		WorkerCount: numWorkers,
		TaskQueue:   make(chan Task, numTasks),
		ResultQueue: make(chan int, numTasks),
	}

	workerPool.Start()

	for i := 1; i <= numTasks; i++ {
		workerPool.TaskQueue <- Task{ID: i, Data: rand.Intn(100)}
	}
	close(workerPool.TaskQueue)

	workerPool.Shutdown()

	fmt.Println("\nProcessed Results:")
	for result := range workerPool.ResultQueue {
		fmt.Println("Result:", result)
	}
}

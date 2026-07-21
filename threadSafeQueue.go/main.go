package main

import (
	"errors"
	"fmt"
	"sync"
)

// https://github.com/blindlobstar/go-interview-problems/tree/main/10-concurrent-queue
type Queue[T any] struct {
	mu    sync.Mutex
	size  int
	Value []T
}

var queueFull = errors.New("Queue is full")

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		mu:    sync.Mutex{},
		size:  size,
		Value: make([]T, 0, size),
	}
}

func (q *Queue[T]) Push(value T) error {
	currLen := len(q.Value)
	if currLen == q.size {
		return queueFull
	}

	q.mu.Lock()
	q.Value = append(q.Value, value)
	q.mu.Unlock()
	return nil
}

func (q *Queue[T]) Pop() T {
	currLen := len(q.Value)
	var zero T
	if currLen == 0 {
		return zero
	}
	q.mu.Lock()
	firstValue := q.Value[0]
	q.Value = q.Value[1:]
	q.mu.Unlock()

	return firstValue
}

func main() {
	queue := NewQueue[int](5)

	queue.Push(1)
	queue.Push(2)
	queue.Push(19)
	queue.Push(43)

	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}

package main

import (
	"fmt"
	"time"
)

func rateLimiter(request chan int, limit int, interval time.Duration) {

	count := 0

	for req := range request {
		if count >= limit {
			timer := time.NewTimer(interval)
			<-timer.C
			count = 0
		}
		fmt.Println("Processing request: ", req)
		count++
	}
}

func main() {

	request := make(chan int, 10)

	go rateLimiter(request, 3, 2*time.Second)

	for i := 0; i <= 10; i++ {
		request <- i
	}
	close(request)

	time.Sleep(10 * time.Second)
}
